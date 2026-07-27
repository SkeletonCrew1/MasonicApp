package main

import (
	"crypto/tls"
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

	message := []byte(fmt.Sprintf("From: %s\r\nSubject: %s\r\n\r\n%s", from, request.Subject, request.Body))

	auth := smtp.PlainAuth("", from, password, host)
	addr := host + ":" + port

	c, err := smtp.Dial(addr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Fail to send mail", http.StatusInternalServerError)
		return
	}
	defer c.Quit()

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	if err = c.StartTLS(tlsConfig); err != nil {
		log.Println(err)
		http.Error(w, "Fail to send mail", http.StatusInternalServerError)
		return
	}

	if err = c.Auth(auth); err != nil {
		log.Println(err)
		http.Error(w, "Fail to send mail", http.StatusInternalServerError)
		return
	}

	if err = c.Mail(from); err != nil {
		log.Println(err)
		http.Error(w, "Fail to send mail", http.StatusInternalServerError)
		return
	}

	for _, addr := range request.Dest {
		if err = c.Rcpt(addr); err != nil {
			log.Println(err)
			http.Error(w, "Fail to send mail", http.StatusInternalServerError)
			return
		}
	}

	wWriter, err := c.Data()
	if err != nil {
		log.Println(err)
		http.Error(w, "Fail to send mail", http.StatusInternalServerError)
		return
	}

	if _, err = wWriter.Write(message); err != nil {
		log.Println(err)
		http.Error(w, "Fail to send mail", http.StatusInternalServerError)
		return
	}

	if err = wWriter.Close(); err != nil {
		log.Println(err)
		http.Error(w, "Fail to send mail", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/sent-mail", sendMailHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
