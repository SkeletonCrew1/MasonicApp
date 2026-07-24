package handlers

import (
	"dailyfunctions/internal/models"
	"database/sql"
)

func GetAllUsers(db *sql.DB) ([]models.User, error) {
	rows, err := db.Query("SELECT login, status, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.Login, &u.Status, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func GetFilteredUsers(db *sql.DB) ([]models.User, error) {
	rows, err := db.Query("SELECT login, status, email FROM users WHERE status IN ('golden', 'silver')")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.Login, &u.Status, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func ResetInquisitors(db *sql.DB) error {
	_, err := db.Exec("UPDATE users SET is_inquisitor = false")
	return err
}

func SetUserAsInquisitor(db *sql.DB, login string) error {
	_, err := db.Exec("UPDATE users SET is_inquisitor = true WHERE login = $1", login)
	return err
}
