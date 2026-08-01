package models

type Sighting struct {
	ID          int     `json:"SightingId"`
	Name        string  `json:"SightingName"`
	Latitude    float64 `json:"SightingLatitude"`
	Longitude   float64 `json:"SightingLongitude"`
	Date        string  `json:"SightingDiscoveryDate"`
	Picture     string  `json:"SightingPicture"`
	Description string  `json:"SightingDescription"`
	SeenCount   int     `json:"SeenCount"`
	HasSeenIt   bool    `json:"HasSeenIt"`
}
