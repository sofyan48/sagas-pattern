package entity

import "time"

// LoggerEventHistory ...
type LoggerEventHistory struct {
	UUID      string          `json:"__uuid"`
	Action    string          `json:"__action"`
	Offset    uint64          `json:"__offset"`
	History   []LoggerHistory `json:"history"`
	Status    string          `json:"status"`
	CreatedAt *time.Time      `json:"created_at"`
	UpdateAt  *time.Time      `json:"update_at"`
}

// LoggerHistory ...
type LoggerHistory struct {
	Data map[string]interface{} `json:"data"`
}
