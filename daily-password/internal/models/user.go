package models

type User struct {
	ID            int    `json:"UserId"`
	Login         string `json:"UserDisplayName"`
	Password      string `json:"UserPassword"`
	Status        string `json:"UserStatus"`
	Email         string `json:"UserEmail"`
	Is_inquisitor bool   `json:"UserIsInquisitor"`
}
