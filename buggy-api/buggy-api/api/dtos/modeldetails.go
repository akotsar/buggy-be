package dtos

import (
	"buggy/internal/data/modeldata"
	"time"
)

// ModelDetails contains details of a car model.
type ModelDetails struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Image       string          `json:"image"`
	Make        string          `json:"make"`
	MakeID      string          `json:"makeId"`
	MakeImage   string          `json:"makeImage"`
	Votes       int             `json:"votes"`
	EngineVol   float64         `json:"engineVol"`
	MaxSpeed    int             `json:"maxSpeed"`
	Comments    []*ModelComment `json:"comments"`
	CanVote     bool            `json:"canVote"`
}

// ModelComment represents a single comment.
type ModelComment struct {
	User       string    `json:"user"`
	DatePosted time.Time `json:"datePosted"`
	Text       string    `json:"text"`
}

// NewModelDetailsFromRecord creates an instance of ModelDetails from a model db record.
func NewModelDetailsFromRecord(record *modeldata.ModelRecord) *ModelDetails {
	return &ModelDetails{
		Name:        record.Name,
		Description: record.Description,
		Image:       record.Image,
		MakeID:      record.GetMakeID(),
		Votes:       record.Votes,
		EngineVol:   record.EngineVol,
		MaxSpeed:    record.MaxSpeed,
	}
}
