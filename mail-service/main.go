package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
)

type MailRequest struct {
	Dest    []string `json:"dest"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

func sendMailHandler(w http.ResponseWriter, r *http.Request) {
	var request MailRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	from := os.Getenv("SENDER")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("SMTPHOST")
	port := os.Getenv("SMTPPORT")

	message := []byte(fmt.Sprintf("Subject: %s\n%s", request.Subject, request.Body))
	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, request.Dest, message)
	if err != nil {
		log.Println(err)
		http.Error(w, "Fail to send mail", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/sent-mail", sendMailHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
