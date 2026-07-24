package models

type User struct {
	ID            int    `json:"id"`
	Login         string `json:"login"`
	Status        string `json:"status"`
	Email         string `json:"email"`
	Is_inquisitor bool   `json:"is_inquisitor"`
}
