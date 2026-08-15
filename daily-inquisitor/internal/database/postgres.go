package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	host := os.Getenv("USERS_DB_HOST")
	port := os.Getenv("USERS_DB_PORT")
	user := os.Getenv("USERS_DB_USER")
	pass := os.Getenv("USERS_DB_PASSWORD")
	name := os.Getenv("USERS_DB_NAME")

	conn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		host, port, user, pass, name,
	)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	return db, nil
}
