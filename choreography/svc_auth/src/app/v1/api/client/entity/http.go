package entity

import "time"

// ClientRequest ...
type ClientRequest struct {
	ClientName   string `json:"client_name"`
	IsFirstParty bool   `json:"is_first_party"`
	RedirectURIs string `json:"redirect_url"`
}

// ClientResponses ...
type ClientResponses struct {
	ID              int64      `gorm:"column:id_client;primary_key" json:"id_client"`
	ClientName      string     `gorm:"column:client_name" json:"client_name"`
	ClientKey       string     `gorm:"column:client_key" json:"client_key"`
	ClientSecret    string     `gorm:"column:client_secret" json:"client_secret"`
	ClientPUblicKey string     `gorm:"column:client_public_key" json:"client_public_key"`
	IsActive        bool       `gorm:"column:is_active" json:"is_active"`
	IsFirtsParty    bool       `gorm:"column:is_first_party" json:"is_first_party"`
	RedirectUrls    string     `gorm:"column:redirect_url" json:"redirect_url"`
	CreatedAt       *time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       *time.Time `gorm:"column:update_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt       *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
