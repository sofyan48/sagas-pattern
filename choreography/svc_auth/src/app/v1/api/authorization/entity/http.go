package entity

// AuhtorizerClientResponse ...
type AuhtorizerClientResponse struct {
	Expire  interface{}                      `json:"expire_at"`
	Issuer  interface{}                      `json:"issuer"`
	Secret  interface{}                      `json:"secret"`
	Session *AuhtorizerClientResponseSession `json:"session"`
}

// AuhtorizerClientResponseSession ...
type AuhtorizerClientResponseSession struct {
	Level string `json:"level"`
}
