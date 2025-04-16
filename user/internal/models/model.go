package models


type User struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Bio      string `json:"bio,omitempty"`
}