package dtos

import "buggy/internal/data/makedata"

// Make represents a car make.
type Make struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Models      ModelList `json:"models"`
}

type ModelList struct {
	TotalPages int         `json:"totalPages"`
	Models     []ModelItem `json:"models"`
}

func NewMakeFromRecord(source *makedata.MakeRecord) *Make {
	if source == nil {
		return nil
	}

	return &Make{
		Name:        source.Name,
		Description: source.Description,
		Image:       source.Image,
	}
}
