package entity

import (
	"time"
)

// UserEvent ...
type UserEvent struct {
	UUID      string                 `json:"__uuid" `
	Action    string                 `json:"__action"`
	Offset    int64                  `json:"__offset"`
	Data      map[string]interface{} `json:"data"`
	Status    string                 `json:"status"`
	CreatedAt *time.Time             `json:"created_at"`
}
