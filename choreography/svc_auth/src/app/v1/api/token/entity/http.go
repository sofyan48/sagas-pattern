package entity

// ClientCredentialRequest ...
type ClientCredentialRequest struct {
	ClientKey    string `json:"client_key"`
	ClientSecret string `json:"client_secret"`
}

type ClientCredentialResponse struct {
	Type        string `json:"token_type"`
	AccessToken string `json:"access_token"`
	Expires     int64  `json:"expires_at"`
}
