package main

type User struct {
	UserEmail     string `json:"UserEmail"`
	UserPassword  string `json:"UserPassword"`
	DailyPassword string `json:"DailyPassword"`
}
type CustomClaims struct {
	username string `json:"username"`
	status   bool   `json:"status"`
	userid   string `json:"userid"`
}
