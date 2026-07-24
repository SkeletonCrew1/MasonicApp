package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("USERS_DB_NAME")

	conn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name,
	)

	var db *sql.DB
	var err error

	// Retry loop to handle the race condition on startup
	for i := 0; i < 15; i++ {
		db, err = sql.Open("postgres", conn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				return db, nil
			}
		}

		log.Printf("Database not ready yet, retrying in 2 seconds... (%v)", err)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("could not connect to database after multiple retries: %v", err)
}
