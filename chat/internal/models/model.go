package models

import "time"

type Chat struct {
	ID      string   `json:"id"`
	Members []string `json:"members"`
}

type Message struct {
	ID        string    `json:"id"`
	ChatID    string    `json:"chat_id"`
	SenderID  string    `json:"sender_id"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}