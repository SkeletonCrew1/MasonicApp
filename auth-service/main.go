package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	_ "runtime/trace"
	"time"

	_ "github.com/golang-jwt/jwt/v5"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

type Login struct {
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}

var fake_name_list = [30]string{
	"James", "Michael", "John", "Robert", "David",
	"William", "Richard", "Joseph", "Thomas", "Christopher",
	"Charles", "Daniel", "Matthew", "Anthony", "Mark",
	"Steven", "Andrew", "Donald", "Joshua", "Paul",
	"Kevin", "Kenneth", "Brian", "Timothy", "Ronald",
	"Jason", "George", "Edward", "Jeffrey", "Jacob"}

var users = map[string]Login{}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/register", register)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/protected", protected)

	log.Fatal(http.ListenAndServe(":8081", enableCORS(mux)))
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		er := http.StatusMethodNotAllowed
		http.Error(w, "Invalid method", er)
		return
	}
	var register_body User
	if err := json.NewDecoder(r.Body).Decode(&register_body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	random_num := randRange(0, 30)

	user_password := register_body.UserPassword
	user_fake_name := fake_name_list[random_num]
	user_rank := "bronze"
	user_email := register_body.UserEmail
	user_is_inqusitor := false
	//user_secret_key=r.FormValue("user_secret_key")

	if len(user_password) < 8 {
		er := http.StatusNotAcceptable
		http.Error(w, "Invalid name or password", er)
		return
	}

	if valid_email(user_email) == false {
		er := http.StatusConflict
		http.Error(w, "incorect email", er)
		return
	}

	hashedPassword, _ := hashPassword(user_password)

	connStr := os.Getenv("AUTH_URL")

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	defer db.Close()
	db, err = sql.Open("postgres", connStr)
	name_exists := UserExists(db, user_fake_name)
	for {
		name_exists = UserExists(db, user_fake_name)
		if name_exists == true {
			random_num = randRange(0, 30)
			user_fake_name = fake_name_list[random_num]

		} else {
			break
		}
	}

	email_exist := EmailExists(db, user_email)
	if email_exist == true {
		er := http.StatusMethodNotAllowed
		http.Error(w, "email registered", er)
		return
	}

	query := `
		INSERT INTO users (UserDisplayName,UserPassword,UserStatus,UserEmail,UserIsInquisitor)
		VALUES ($1, $2, $3, $4, $5);
	`
	_, err = db.Exec(query, user_fake_name, hashedPassword, user_rank, user_email, user_is_inqusitor)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Fprintf(w, "User %s registered successfully", user_fake_name)

}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		er := http.StatusMethodNotAllowed
		http.Error(w, "Invalid method", er)
		return
	}
	w.Header().Set("Content-type", "application")

	connStr := os.Getenv("AUTH_URL")

	var secretKey = []byte(os.Getenv("SECRET_KEY"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Fprint(w, "can't connect to database")
		panic(err)
	}
	var login_body User
	if err := json.NewDecoder(r.Body).Decode(&login_body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user_password := login_body.UserPassword
	user_email := login_body.UserEmail
	hashedPassword, _ := hashPassword(user_password)
	email_exist := EmailExists(db, user_email)
	user_fake_name := GetUserDisplayName(db, user_email)
	user_status := GetUserStatus(db, user_email)
	user_id := GetUserId(db, user_email)

	if email_exist != true {
		er := http.StatusMethodNotAllowed
		http.Error(w, "Invalid email", er)
		return
	}

	if !checkPasswordHash(user_password, hashedPassword) {
		er := http.StatusUnauthorized
		http.Error(w, "Invalid   password", er)

		return
	}

	tokenString, err := createToken(secretKey, user_fake_name, user_status, user_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer db.Close()
	http.SetCookie(w, &http.Cookie{
		Name:     "JWT",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 2),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		Secure:   true,
	})

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Login successfull")
}

func logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		er := http.StatusMethodNotAllowed
		http.Error(w, "Invalid method", er)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "JWT",
		Value:    "",
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	})
	fmt.Fprint(w, "Logout successfull")

}

func protected(w http.ResponseWriter, r *http.Request) {

}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_SERVICE_URL"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
