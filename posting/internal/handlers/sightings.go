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
)

func AddSighting(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		_, err := db.Exec(`
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
			sightings = append(sightings, s)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(sightings)
	}
}

func GetSightingByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		var s models.Sighting
		err := db.QueryRow("SELECT SightingId, SightingName, SightingLatitude, SightingLongitude, SightingDiscoveryDate, SightingPicture, SightingDescription FROM Sightings WHERE SightingId = $1", id).
			Scan(&s.ID, &s.Name, &s.Latitude, &s.Longitude, &s.Date, &s.Picture, &s.Description)

		if err != nil {
			http.Error(w, "Sighting not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(s)
	}
}
