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

	"github.com/golang-jwt/jwt/v5"
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
	daily_password := register_body.DailyPassword
	hashedPassword, _ := hashPassword(user_password)

	connStr := os.Getenv("AUTH_URL")

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if len(user_password) < 8 {
		er := http.StatusConflict
		http.Error(w, "Password has to be at least 8 characters long", er)
		return
	}

	if valid_email(user_email) == false {
		er := http.StatusConflict
		http.Error(w, "Incorect email", er)
		return
	}

	if !checkDailyPassword(db, daily_password) {
		er := http.StatusConflict
		http.Error(w, "Wrong daily password", er)
		return
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
		er := http.StatusBadRequest
		http.Error(w, "Email registered", er)
		return
	}

	invited := userInvited(db, user_email)
	if invited != true {
		er := http.StatusUnauthorized
		http.Error(w, "You are not invited", er)
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
		log.Println("Invalid token", err)

		panic(err)
	}
	var login_body User
	if err := json.NewDecoder(r.Body).Decode(&login_body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user_password := login_body.UserPassword
	user_email := login_body.UserEmail
	daily_password := login_body.DailyPassword

	hashedPassword, _ := hashPassword(user_password)
	email_exist := EmailExists(db, user_email)
	user_exluded := userExluded(db, user_email)

	user_fake_name := GetUserDisplayName(db, user_email)
	user_status := GetUserStatus(db, user_email)
	user_id := GetUserId(db, user_email)
	is_inquisitor := getUserInquisitor(db, user_email)

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
	if user_exluded {
		er := http.StatusUnauthorized
		http.Error(w, "User has been excluded", er)

		return
	}
	if !checkDailyPassword(db, daily_password) {
		er := http.StatusConflict
		http.Error(w, "Wrong daily password", er)
		return
	}
	tokenString, err := createToken(secretKey, user_fake_name, user_status, user_id, is_inquisitor)
	if err != nil {

		er := http.StatusInternalServerError
		http.Error(w, "Invalid method", er)
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
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		Secure:   true,
	})
	log.Println("Logout successfull")
	w.WriteHeader(http.StatusOK)
}

func protected(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		er := http.StatusMethodNotAllowed
		http.Error(w, "Invalid method", er)
		return
	}

	var secretKey = []byte(os.Getenv("SECRET_KEY"))
	var claims = jwt.MapClaims{}
	JWT_value := GetJWTValue(w, r)
	token, err := jwt.ParseWithClaims(JWT_value, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil // Ensure 'secret' is your HS256 key
	})

	if err != nil || !token.Valid {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		log.Println("Invalid token", err)
		return
	}

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
