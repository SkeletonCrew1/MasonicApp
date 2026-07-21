package main

import (
	"math/rand/v2"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func valid_email(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
