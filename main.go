package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"os"
	_ "runtime/trace"

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
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/protected", protected)
	http.ListenAndServe(":8080", nil)
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		er := http.StatusMethodNotAllowed
		http.Error(w, "Invalid method", er)
		return
	}

	random_num := randRange(0, 30)

	user_password := r.FormValue("user_password")
	user_fake_name := fake_name_list[random_num]
	user_rank := "bronze"
	user_email := r.FormValue("user_email")
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
	users[user_fake_name] = Login{
		HashedPassword: hashedPassword,
	}
	//postgres://postgres:mysecretpassword@users_db:5432/auth_service?sslmode=disable
	connStr := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		er := http.StatusMethodNotAllowed
		http.Error(w, "error4", er)
		return
	}

	check_name_query := `
		SELECT count(1) > 0
		FROM users
		WHERE UserFakename = $1;
	`
	defer db.Close()
	db, err = sql.Open("postgres", connStr)
	for {
		rows, err := db.Query(check_name_query, user_fake_name)
		if err != nil && err != sql.ErrNoRows {
			er := http.StatusMethodNotAllowed
			fmt.Fprintf(w, " %s ", err)
			http.Error(w, "error3", er)
			return
		}
		if err == sql.ErrNoRows {
			random_num = randRange(0, 30)
			user_fake_name = fake_name_list[random_num]

		} else {
			break
		}
		rows.Close()
	}

	check_email_query := `
		SELECT count(1) > 0
		FROM users
		WHERE UserEmail = $1;
	`
	rows, err := db.Query(check_email_query, user_fake_name)
	if err != nil && err != sql.ErrNoRows {
		er := http.StatusMethodNotAllowed
		http.Error(w, "error2", er)
		return
	}
	if err == sql.ErrNoRows {
		er := http.StatusMethodNotAllowed
		http.Error(w, "Email already registered", er)
		return
	}
	rows.Close()

	query := `
		INSERT INTO users (UserFakename,UserPassword,UserStatus,UserEmail)
		VALUES ($1, $2, $3, $4);
	`
	_, err = db.Exec(query, user_fake_name, hashedPassword, user_rank, user_email)
	if err != nil {
		er := http.StatusMethodNotAllowed
		http.Error(w, "error1", er)
		return
	}
	rows.Close()
	defer db.Close()
	fmt.Fprintf(w, "User %s registered successfully", user_fake_name)

}

func login(w http.ResponseWriter, r *http.Request) {}

func logout(w http.ResponseWriter, r *http.Request) {}

func protected(w http.ResponseWriter, r *http.Request) {}
