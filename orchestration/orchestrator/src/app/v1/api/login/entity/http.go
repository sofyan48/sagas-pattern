package entity

import "time"

// UserRequest ...
type UserRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"handphone"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Province    string `json:"province"`
	District    string `json:"district"`
}

// UserResponses ...
type UserResponses struct {
	UUID      string     `json:"uuid"`
	CreatedAt *time.Time `json:"created_at"`
	Event     *UserEvent `json:"event"`
}
