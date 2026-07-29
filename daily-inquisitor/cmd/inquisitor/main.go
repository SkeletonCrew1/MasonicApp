package main

import (
	"bytes"
	"dailyfunctions/internal/database"
	"dailyfunctions/internal/handlers"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

type EmailPayload struct {
	Dest    []string `json:"dest"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

func runInquisitorTask() {
	log.Println("Running scheduled inquisitor selection task...")

	db, err := database.Connect()
	if err != nil {
		log.Println("Database connection error:", err)
		return
	}
	defer db.Close()

	err = handlers.ResetInquisitors(db)
	if err != nil {
		log.Printf("Failed to reset previous inquisitors: %v", err)
		return
	}

	users, err := handlers.GetFilteredUsers(db)
	if err != nil {
		log.Printf("Failed to fetch filtered users: %v", err)
		return
	}

	if len(users) == 0 {
		log.Println("No eligible golden or silver users found.")
		return
	}

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	randomIndex := rng.Intn(len(users))
	selectedUser := users[randomIndex]

	log.Printf("Selected new inquisitor: %s (%s)", selectedUser.Login, selectedUser.Email)

	err = handlers.SetUserAsInquisitor(db, selectedUser.Login)
	if err != nil {
		log.Printf("Failed to update user inquisitor status: %v", err)
		return
	}

	payload := EmailPayload{
		Dest:    []string{selectedUser.Email},
		Subject: "Congratulations!",
		Body:    "Congratulations, you are an inquisitor today!",
	}

	emailServiceEndpoint := os.Getenv("MAIL_SERVICE_URL")

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON payload: %v", err)
		return
	}

	resp, err := http.Post(emailServiceEndpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Failed to send request to email service: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Email service returned non-ok status: %d", resp.StatusCode)
		return
	}

	log.Println("Inquisitor successfully chosen and notified!")
}

func main() {
	runInquisitorTask()

	s := cron.New()

	_, err := s.AddFunc("0 0 * * *", runInquisitorTask)
	if err != nil {
		log.Fatalf("Failed to add cron job: %v", err)
	}

	s.Start()
	log.Println("Inquisitor service cron scheduler started. Running daily...")

	select {}
}
