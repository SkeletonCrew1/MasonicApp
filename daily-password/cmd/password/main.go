package main

import (
	"bytes"
	"dailyfunctions/internal/database"
	"dailyfunctions/internal/handlers"
	"encoding/json"
	"fmt"
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

func GenerateRandomPassword(passwordLength int) string {
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	specialChar := "!@#$%^&*()_-+={}[/?]"

	password := ""

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	for n := 0; n < passwordLength; n++ {
		randNum := rng.Intn(4)

		switch randNum {
		case 0:
			randCharNum := rng.Intn(len(lowerCase))
			password += string(lowerCase[randCharNum])
		case 1:
			randCharNum := rng.Intn(len(upperCase))
			password += string(upperCase[randCharNum])
		case 2:
			randCharNum := rng.Intn(len(numbers))
			password += string(numbers[randCharNum])
		case 3:
			randCharNum := rng.Intn(len(specialChar))
			password += string(specialChar[randCharNum])
		}
	}

	return password
}

func runPasswordTask() {
	log.Println("Running scheduled password and email task...")

	db, err := database.Connect()
	if err != nil {
		log.Println("Database connection error:", err)
		return
	}
	defer db.Close()

	users, err := handlers.GetAllUsers(db)
	if err != nil {
		log.Printf("Failed to fetch users: %v", err)
		return
	}

	if len(users) == 0 {
		log.Println("No users found to email.")
		return
	}

	var emails []string
	for _, user := range users {
		emails = append(emails, user.Email)
	}

	password := GenerateRandomPassword(20)

	payload := EmailPayload{
		Dest:    emails,
		Subject: "New Password",
		Body:    fmt.Sprintf("Greetings! New password is: %s", password),
	}

	emailServiceEndpoint := os.Getenv("MAIL_SERVICE_URL")

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("failed to marshal JSON payload: %v", err)
		return
	}

	resp, err := http.Post(emailServiceEndpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Failed to send request: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("email service returned non-ok status: %d", resp.StatusCode)
		return
	}

	log.Println("Password emails sent successfully!")
}

func main() {
	runPasswordTask()

	s := cron.New()

	_, err := s.AddFunc("0 0 * * *", runPasswordTask)
	if err != nil {
		log.Fatalf("Failed to add cron job: %v", err)
	}

	s.Start()
	log.Println("Password service cron scheduler started. Running daily...")

	select {}
}
