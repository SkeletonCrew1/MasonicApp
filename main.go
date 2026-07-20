package main

import (
	"database/sql"
	"fmt"
	"math/rand/v2"
	"net/http"
	"net/mail"
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

	user_name := r.FormValue("user_name")
	user_password := r.FormValue("user_password")
	user_fake_name := fake_name_list[random_num]
	user_rank := "bronze"
	user_email := r.FormValue("user_email")
	//user_secret_key=r.FormValue("user_secret_key")

	if len(user_name) < 2 || len(user_password) < 8 {
		er := http.StatusNotAcceptable
		http.Error(w, "Invalid name or password", er)
		return
	}

	if _, ok := users[user_name]; ok {
		er := http.StatusConflict
		http.Error(w, user_name+" already exists", er)
		return
	}

	if valid_email(user_email) == false {
		er := http.StatusConflict
		http.Error(w, "incorect email", er)
		return
	}

	hashedPassword, _ := hashPassword(user_password)
	users[user_name] = Login{
		HashedPassword: hashedPassword,
	}

	for {
		if _, ok := users[user_fake_name]; ok {
			random_num = randRange(0, 30)
			user_fake_name = fake_name_list[random_num]
		} else {
			break
		}
	}
	connStr := os.Getenv("DATABASE_URL")
	print(connStr)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//if err = db.Ping(); err != nil {
	//	log.Println("DB Ping Failed")
	//	log.Fatal(err)
	//}

	query := `
		INSERT INTO users (UserRealName,UserFakename,UserPassword,UserStatus,UserEmail)
		VALUES ($1, $2, $3, $4, $5);
	`
	_, err = db.Exec(query, user_name, user_fake_name, hashedPassword, user_rank, user_email)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Fprintf(w, "User %s registered successfully", user_name)

}

func login(w http.ResponseWriter, r *http.Request) {}

func logout(w http.ResponseWriter, r *http.Request) {}

func protected(w http.ResponseWriter, r *http.Request) {}

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func valid_email(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
