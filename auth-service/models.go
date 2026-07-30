package main

type User struct {
	UserEmail     string `json:"UserEmail"`
	UserPassword  string `json:"UserPassword"`
	DailyPassword string `json:"DailyPassword"`
}
