package dtos

import "buggy/internal/data/modeldata"

// ModelItem represents a car model.
type ModelItem struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Image         string   `json:"image"`
	Make          string   `json:"make"`
	MakeID        string   `json:"makeId"`
	MakeImage     string   `json:"makeImage"`
	Votes         int      `json:"votes"`
	Rank          int      `json:"rank"`
	EngineVol     float64  `json:"engineVol"`
	Comments      []string `json:"comments"`
	TotalComments int      `json:"totalComments"`
}

func NewModelItemFromRecord(source *modeldata.ModelRecord) *ModelItem {
	if source == nil {
		return nil
	}

	return &ModelItem{
		ID:        source.GetIDFromTypeAndID(),
		Name:      source.Name,
		Image:     source.Image,
		MakeID:    source.GetMakeID(),
		Votes:     source.Votes,
		EngineVol: source.EngineVol,
	}
}
