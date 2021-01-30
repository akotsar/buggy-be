package dtos

import "time"

// Comment represents a user comment.
type Comment struct {
	User       string    `json:"user"`
	DatePosted time.Time `json:"datePosted"`
	Text       string    `json:"text"`
}
