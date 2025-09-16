package models

import "time"

type User struct {
	ID                    int
	CardID                string
	FullName              string
	Email                 string
	Phone                 string
	Balance               float32
	Status                string
	CurrentTripTerminalID *int64
	RegisteredAt          time.Time
	UpdatedAt             time.Time
}
