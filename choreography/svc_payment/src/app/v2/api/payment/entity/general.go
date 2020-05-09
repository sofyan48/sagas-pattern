package entity

// Pagination ..
type Pagination struct {
	Limit int `json:"limit" form:"limit"`
	Page  int `json:"page" form:"page"`
}
