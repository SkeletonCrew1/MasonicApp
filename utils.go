package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand/v2"
	"net/mail"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func GetValue(db *sql.DB, column string, email string) string {

	var value string
	rows, err := db.Query("SELECT $1 FROM users WHERE UserEmail = $2;", column, email)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&value); err != nil {
			log.Fatal(err)
		}
		fmt.Println(value)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return value

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

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func createToken(secretKey []byte, username string, email string, status string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"email":    email,
			"status":   status,
			"exp":      time.Now().Add(time.Hour * 24).Unix()})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func verifyToken(tokenString string, secretKey []byte) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("Invalid token")
	}

	return nil
}
