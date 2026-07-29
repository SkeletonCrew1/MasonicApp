package handlers

import (
	"database/sql"
	"fmt"
	"time"
)

func UpdateDailyCode(db *sql.DB, newCode string, emailSent bool) error {
	if !emailSent {
		return fmt.Errorf("Mailing service did not succeed!")
	}

	today := time.Now().Format("2006-01-02")

	_, err := db.Exec(`INSERT INTO Dailycode (DailyCode, CodeDate) VALUES ($1, $2)`, newCode, today)
	return err
}
