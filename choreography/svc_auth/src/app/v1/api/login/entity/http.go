package entity

type LoginResponse struct {
	Type        string `json:"token_type"`
	AccessToken string `json:"access_token"`
	Expires     int64  `json:"expires_at"`
	Refresh     string `json:"refresh_token"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
