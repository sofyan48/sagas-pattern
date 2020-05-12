package entity

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserModelSession struct {
	Login   *LoginDataModels   `json:"login"`
	Session *SessionDataModels `json:"session"`
}

type SessionDataModels struct {
	ID          uint64     `json:"id_user"`
	UUID        string     `json:"uuid"`
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

type LoginDataModels struct {
	IDUser    string     `json:"id_user"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	IDRoles   string     `json:"id_roles"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"update_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Claims struct {
	Session string `json:"session"`
	jwt.StandardClaims
}

type ClientCredentialResponse struct {
	Type         string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	Expires      int64  `json:"expires_at"`
	RefreshToken string `json:"refresh_token"`
}
