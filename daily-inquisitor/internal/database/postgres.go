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
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	name := os.Getenv("MAIN_DB_NAME")

	conn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name,
	)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	maxRetries := 15
	for i := 1; i <= maxRetries; i++ {
		err = db.Ping()
		if err == nil {
			log.Println("Successfully connected to the database!")
			return db, nil
		}

		log.Printf("Database not ready yet, retrying in 2 seconds... (%v) (attempt %d/%d)", err, i, maxRetries)
		time.Sleep(2 * time.Second)
	}

	db.Close()
	return nil, fmt.Errorf("could not connect to database after multiple retries: %v", err)
}
