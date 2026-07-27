package handlers

import (
	"dailyfunctions/internal/models"
	"database/sql"
)

func GetAllUsers(db *sql.DB) ([]models.User, error) {
	rows, err := db.Query(`SELECT UserId, UserDisplayName, UserPassword, UserStatus, UserEmail, UserIsInquisitor FROM "users"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Login, &u.Password, &u.Status, &u.Email, &u.Is_inquisitor); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func GetFilteredUsers(db *sql.DB) ([]models.User, error) {
	rows, err := db.Query(`SELECT UserId, UserDisplayName, UserPassword, UserStatus, UserEmail, UserIsInquisitor FROM "users" WHERE UserStatus IN ('golden', 'silver')`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Login, &u.Password, &u.Status, &u.Email, &u.Is_inquisitor); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func ResetInquisitors(db *sql.DB) error {
	_, err := db.Exec(`UPDATE "users" SET UserIsInquisitor = false`)
	return err
}

func SetUserAsInquisitor(db *sql.DB, login string) error {
	_, err := db.Exec(`UPDATE "users" SET UserIsInquisitor = true WHERE UserDisplayName = $1`, login)
	return err
}
