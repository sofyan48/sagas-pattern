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
	SiteProfil  string     `json:"site_profil"`
	PhoneNumber string     `json:"handphone"`
	Address     string     `json:"address"`
	City        string     `json:"city"`
	Province    string     `json:"province"`
	District    string     `json:"district"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

// StateFullFormatKafka ...
type StateFullFormatKafka struct {
	UUID      string            `json:"__uuid" bson:"__uuid"`
	Action    string            `json:"__action" bson:"__action"`
	Data      map[string]string `json:"data" bson:"data"`
	CreatedAt *time.Time        `json:"created_at" bson:"created_at"`
}
