package main

import (
	"database/sql"
	"log"
	"math/rand/v2"
	"net/http"
	"net/mail"

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

func GetUserStatus(db *sql.DB, email string) string {

	var value string
	rows, err := db.Query("SELECT UserStatus FROM users WHERE UserEmail = $1;", email)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&value); err != nil {
			log.Fatal(err)
		}

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return value
}

func GetUserDisplayName(db *sql.DB, email string) string {
	var value string
	rows, err := db.Query("SELECT UserDisplayName FROM users WHERE UserEmail = $1;", email)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&value); err != nil {
			log.Fatal(err)
		}

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return value
}

func GetUserId(db *sql.DB, email string) string {

	var value string
	rows, err := db.Query("SELECT UserIsInquisitor FROM users WHERE UserEmail = $1;", email)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&value); err != nil {
			log.Fatal(err)
		}

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return value
}

func getUserInquisitor(db *sql.DB, email string) string {

	var value string
	rows, err := db.Query("SELECT UserIsInquisitor FROM users WHERE UserEmail = $1;", email)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&value); err != nil {
			log.Fatal(err)
		}

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return value
}

func userInvited(db *sql.DB, email string) bool {
	row := db.QueryRow("select InvitedEmail from WhiteList where InvitedEmail= $1", email)
	temp := ""
	row.Scan(&temp)
	if temp != "" {
		return true
	}
	return false
}

func userExluded(db *sql.DB, email string) bool {
	row := db.QueryRow("select BannedEmail from Blacklist where BannedEmail= $1", email)
	temp := ""
	row.Scan(&temp)
	if temp != "" {
		return true
	}
	return false
}

func UserExists(db *sql.DB, username string) bool {
	sqlStmt := `SELECT UserDisplayName FROM users WHERE UserDisplayName = $1`
	err := db.QueryRow(sqlStmt, username).Scan(&username)
	if err != nil {
		if err != sql.ErrNoRows {

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

func createToken(secretKey []byte, username string, status string, userid string, is_inquisitor string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username":      username,
			"status":        status,
			"userid":        userid,
			"is_inquisitor": is_inquisitor,
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetJWTValue(w http.ResponseWriter, r *http.Request) string {
	cookie, err := r.Cookie("JWT")
	if err != nil {
		return ""
	}

	return cookie.Value
}

func GetDailyPasscode(db *sql.DB) string {

	var value string
	rows, err := db.Query("SELECT DailyCode FROM dailycode ORDER BY CodeId DESC LIMIT 1;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&value); err != nil {
			log.Fatal(err)
		}

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return value
}

func checkDailyPassword(db *sql.DB, input_pass string) bool {
	daily_pass := GetDailyPasscode(db)

	if daily_pass == input_pass {
		return true
	} else {
		return false
	}

}
