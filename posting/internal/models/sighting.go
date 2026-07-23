package models

type Sighting struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Date        string  `json:"date"`
	Picture     string  `json:"picture"`
	Description string  `json:"description"`
}
