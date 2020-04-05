package entity

import "time"

// HealtResponse ...
type HealtResponse struct {
	Status    string     `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
}
