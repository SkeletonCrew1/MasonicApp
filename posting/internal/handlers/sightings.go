package handlers

import (
	"backend/internal/models"
	"bytes"
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/golang-jwt/jwt/v5"
)

func getS3PresignClient() (*s3.PresignClient, string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, "", err
	}
	client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(client)
	bucket := os.Getenv("S3_BUCKET_NAME")
	return presignClient, bucket, nil
}

func generatePresignedURL(objectKey string) (string, error) {
	if objectKey == "" {
		return "", nil
	}
	presignClient, bucket, err := getS3PresignClient()
	if err != nil {
		return "", err
	}

	req, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &objectKey,
	}, s3.WithPresignExpires(15*time.Minute))

	if err != nil {
		return "", err
	}
	return req.URL, nil
}

func AddSighting(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status, err := getUserStatusFromRequest(r)
		if err != nil || (status != "silver" && status != "golden" && status != "gold") {
			log.Println("Auth error in AddSighting:", err, "Status:", status)
			http.Error(w, "Forbidden: insufficient permissions to create posts", http.StatusForbidden)
			return
		}

		var sighting models.Sighting
		if err := json.NewDecoder(r.Body).Decode(&sighting); err != nil {
			log.Println("JSON decode error:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var s3ObjectKey string
		if sighting.Picture != "" {
			idx := strings.Index(sighting.Picture, ",")
			if idx != -1 {
				sighting.Picture = sighting.Picture[idx+1:]
			}

			data, err := base64.StdEncoding.DecodeString(sighting.Picture)
			if err != nil {
				log.Println("Base64 decode error:", err)
				http.Error(w, "Invalid image data", http.StatusBadRequest)
				return
			}

			filename := fmt.Sprintf("%d.png", time.Now().UnixNano())
			s3ObjectKey = "uploads/" + filename

			cfg, err := config.LoadDefaultConfig(context.TODO())
			if err != nil {
				log.Println("AWS LoadDefaultConfig error:", err)
				http.Error(w, "Failed to connect to AWS", http.StatusInternalServerError)
				return
			}
			s3Client := s3.NewFromConfig(cfg)
			bucket := os.Getenv("S3_BUCKET_NAME")

			_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
				Bucket:      aws.String(bucket),
				Key:         aws.String(s3ObjectKey),
				Body:        bytes.NewReader(data),
				ContentType: aws.String("image/png"),
			})
			if err != nil {
				log.Println("S3 PutObject error:", err)
				http.Error(w, "Failed to upload image to S3", http.StatusInternalServerError)
				return
			}
		}

		_, err = db.Exec(`
            INSERT INTO sightings (SightingName, SightingLatitude, SightingLongitude, SightingDiscoveryDate, SightingPicture, SightingDescription)
            VALUES($1, $2, $3, $4, $5, $6)
        `, sighting.Name, sighting.Latitude, sighting.Longitude, sighting.Date, s3ObjectKey, sighting.Description)

		if err != nil {
			log.Println("Database insert error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func GetSightings(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status, err := getUserStatusFromRequest(r)
		if err != nil || (status != "bronze" && status != "silver" && status != "golden" && status != "gold") {
			http.Error(w, "Forbidden: insufficient permissions to view posts", http.StatusForbidden)
			return
		}

		currentUsername, _ := getUserNameFromRequest(r)

		rows, err := db.Query("SELECT SightingId, SightingName, SightingLatitude, SightingLongitude, SightingDiscoveryDate, SightingPicture, SightingDescription FROM Sightings")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var sightings []models.Sighting
		for rows.Next() {
			var s models.Sighting
			var s3Key string
			if err := rows.Scan(&s.ID, &s.Name, &s.Latitude, &s.Longitude, &s.Date, &s3Key, &s.Description); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if s3Key != "" {
				presignedURL, err := generatePresignedURL(s3Key)
				if err == nil {
					s.Picture = presignedURL
				}
			}

			db.QueryRow("SELECT COUNT(*) FROM SeenToo WHERE post = $1", fmt.Sprintf("%d", s.ID)).Scan(&s.SeenCount)

			if currentUsername != "" {
				var count int
				db.QueryRow("SELECT COUNT(*) FROM SeenToo WHERE post = $1 AND user_identifier = $2", fmt.Sprintf("%d", s.ID), currentUsername).Scan(&count)
				s.HasSeenIt = count > 0
			}

			sightings = append(sightings, s)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(sightings)
	}
}

func GetSightingByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status, err := getUserStatusFromRequest(r)
		if err != nil || (status != "bronze" && status != "silver" && status != "golden" && status != "gold") {
			http.Error(w, "You are not allowed to see our map!", http.StatusForbidden)
			return
		}

		id := r.URL.Query().Get("id")
		var s models.Sighting
		var s3Key string
		err = db.QueryRow("SELECT SightingId, SightingName, SightingLatitude, SightingLongitude, SightingDiscoveryDate, SightingPicture, SightingDescription FROM Sightings WHERE SightingId = $1", id).
			Scan(&s.ID, &s.Name, &s.Latitude, &s.Longitude, &s.Date, &s3Key, &s.Description)

		if err != nil {
			http.Error(w, "Sighting not found", http.StatusNotFound)
			return
		}

		if s3Key != "" {
			presignedURL, err := generatePresignedURL(s3Key)
			if err == nil {
				s.Picture = presignedURL
			}
		}

		db.QueryRow("SELECT COUNT(*) FROM SeenToo WHERE post = $1", id).Scan(&s.SeenCount)

		currentUsername, err := getUserNameFromRequest(r)
		if err == nil && currentUsername != "" {
			var count int
			db.QueryRow("SELECT COUNT(*) FROM SeenToo WHERE post = $1 AND user_identifier = $2", id, currentUsername).Scan(&count)
			s.HasSeenIt = count > 0
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(s)
	}
}

func awsString(v string) *string {
	return &v
}

func getUserStatusFromRequest(r *http.Request) (string, error) {
	cookie, err := r.Cookie("JWT")
	if err != nil {
		return "", fmt.Errorf("missing token cookie")
	}

	secretKey := []byte(os.Getenv("SECRET_KEY"))
	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}

	status, ok := claims["status"].(string)
	if !ok {
		return "", fmt.Errorf("status claim missing")
	}

	return status, nil
}

func getUserNameFromRequest(r *http.Request) (string, error) {
	cookie, err := r.Cookie("JWT")
	if err != nil {
		return "", fmt.Errorf("missing token cookie")
	}

	secretKey := []byte(os.Getenv("SECRET_KEY"))
	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", fmt.Errorf("username claim missing")
	}

	return username, nil
}

func AddSeenToo(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		username, err := getUserNameFromRequest(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		var payload struct {
			PostID string `json:"postId"`
		}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil || payload.PostID == "" {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		_, err = db.Exec(`
			INSERT INTO SeenToo (post, user_identifier)
			VALUES ($1, $2)
			ON CONFLICT (post, user_identifier) DO NOTHING
		`, payload.PostID, username)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Successfully recorded approval")
	}
}

func GetUserStatusHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status, err := getUserStatusFromRequest(r)
		if err != nil {
			status = "bronze"
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": status})
	}
}
