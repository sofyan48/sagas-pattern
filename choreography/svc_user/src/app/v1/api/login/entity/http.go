package entity

import "time"

// UserLoginRequest ...
type UserLoginRequest struct {
	IDUser   string `json:"id_user"`
	Username string `json:"username"`
	Password string `json:"password"`
	IDRoles  string `json:"id_roles"`
}

// UserLoginResponses ...
type UserLoginResponses struct {
	UUID      string          `json:"uuid"`
	CreatedAt *time.Time      `json:"created_at"`
	Event     *UserLoginEvent `json:"event"`
}

// LoginResponse ...
type LoginResponse struct {
	IDUser    string     `json:"id_user"`
	Username  string     `json:"username"`
	IDRoles   string     `json:"id_roles"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"update_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// Pagination ...
type Pagination struct {
	Limit int `form:"limit" json:"limit"`
	Page  int `form:"page" json:"page"`
}

// GetByUsernameRequest ...
type GetByUsernameRequest struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

// GetByIDRequest ...
type GetByIDRequest struct {
	ID string `form:"id_user" json:"id_user"`
}

type SessionReponse struct {
	Login   interface{} `json:"login"`
	Session interface{} `json:"session"`
}
