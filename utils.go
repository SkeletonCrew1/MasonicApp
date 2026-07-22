package main

import (
	"database/sql"
	"log"
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

func UserExists(db *sql.DB, username string) bool {
	sqlStmt := `SELECT UserFakename FROM users WHERE UserFakename = $1`
	err := db.QueryRow(sqlStmt, username).Scan(&username)
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened! you should change your function return
			// to "(bool, error)" and return "false, err" here
			log.Print(err)
		}

		return false
	}

	return true
}

func EmailExists(db *sql.DB, email string) bool {
	sqlStmt := `SELECT UserEmail FROM users WHERE UserEmail = $1`
	err := db.QueryRow(sqlStmt, email).Scan(&email)
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened! you should change your function return
			// to "(bool, error)" and return "false, err" here
			log.Print(err)
		}

		return false
	}

	return true
}
