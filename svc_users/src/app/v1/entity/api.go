package entity

import "time"

// UsersPayload Mapping
type UsersPayload struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"handphone"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Province    string `json:"province"`
	District    string `json:"district"`
}

// UsersResponse Mapping
type UsersResponse struct {
	ID          uint       `json:"id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Email       string     `json:"email"`
	PhoneNumber string     `json:"handphone"`
	Address     string     `json:"address"`
	City        string     `json:"city"`
	Province    string     `json:"province"`
	District    string     `json:"district"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
