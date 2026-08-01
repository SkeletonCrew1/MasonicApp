package handlers

import (
	"backend/internal/models"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func AddSighting(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status, err := getUserStatusFromRequest(r)
		if err != nil || (status != "silver" && status != "golden" && status != "gold") {
			http.Error(w, "Forbidden: insufficient permissions to create posts", http.StatusForbidden)
			return
		}

		var sighting models.Sighting
		if err := json.NewDecoder(r.Body).Decode(&sighting); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if sighting.Picture != "" {
			idx := strings.Index(sighting.Picture, ",")
			if idx != -1 {
				sighting.Picture = sighting.Picture[idx+1:]
			}

			data, err := base64.StdEncoding.DecodeString(sighting.Picture)
			if err != nil {
				http.Error(w, "Invalid image data", http.StatusBadRequest)
				return
			}

			filename := fmt.Sprintf("%d.png", time.Now().UnixNano())
			filePath := "uploads/" + filename
			if err := os.WriteFile(filePath, data, 0644); err != nil {
				http.Error(w, "Failed to save image", http.StatusInternalServerError)
				return
			}
			sighting.Picture = "/uploads/" + filename
		}

		_, err = db.Exec(`
			INSERT INTO sightings (SightingName, SightingLatitude, SightingLongitude, SightingDiscoveryDate, SightingPicture, SightingDescription)
			VALUES($1, $2, $3, $4, $5, $6)
		`, sighting.Name, sighting.Latitude, sighting.Longitude, sighting.Date, sighting.Picture, sighting.Description)

		if err != nil {
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
			if err := rows.Scan(&s.ID, &s.Name, &s.Latitude, &s.Longitude, &s.Date, &s.Picture, &s.Description); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
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
		err = db.QueryRow("SELECT SightingId, SightingName, SightingLatitude, SightingLongitude, SightingDiscoveryDate, SightingPicture, SightingDescription FROM Sightings WHERE SightingId = $1", id).
			Scan(&s.ID, &s.Name, &s.Latitude, &s.Longitude, &s.Date, &s.Picture, &s.Description)

		if err != nil {
			http.Error(w, "Sighting not found", http.StatusNotFound)
			return
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
