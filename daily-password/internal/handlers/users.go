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
