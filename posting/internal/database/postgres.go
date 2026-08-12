package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	host := os.Getenv("MAP_DB_HOST")
	port := os.Getenv("MAP_DB_PORT")
	user := os.Getenv("MAP_DB_USER")
	pass := os.Getenv("MAP_DB_PASSWORD")
	name := os.Getenv("MAP_DB_NAME")

	conn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		host, port, user, pass, name,
	)

	return sql.Open("postgres", conn)
}
