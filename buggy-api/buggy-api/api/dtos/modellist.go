package dtos

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

// ModelList represents a collection of models
type ModelList struct {
	TotalPages int          `json:"totalPages"`
	Models     []*ModelItem `json:"models"`
}

// GetModelsPage returns a single page of models.
func GetModelsPage(models []*ModelItem, pageIndex int, pageSize int, orderBy string) ModelList {
	// Sorting by Votes first
	sort.Slice(models, func(i, j int) bool {
		return models[i].Votes > models[j].Votes
	})
	for i, model := range models {
		model.Rank = i + 1
	}

	switch orderBy {
	case "make":
		sort.Slice(models, func(i, j int) bool {
			return models[i].Make < models[j].Make
		})
	case "name":
		sort.Slice(models, func(i, j int) bool {
			return models[i].Name < models[j].Name
		})
	case "votes":
		sort.Slice(models, func(i, j int) bool {
			return models[i].Votes > models[j].Votes
		})
	case "rank":
		sort.Slice(models, func(i, j int) bool {
			return fmt.Sprintf("%d", models[i].Rank) < fmt.Sprintf("%d", models[j].Rank)
		})
	case "engine":
		sort.Slice(models, func(i, j int) bool {
			return models[i].EngineVol < models[j].EngineVol
		})
	case "random":
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(models), func(i, j int) { models[i], models[j] = models[j], models[i] })
	}

	skip := (pageIndex - 1) * pageSize
	if skip > len(models) {
		skip = len(models)
	}

	take := pageSize
	if take+skip >= len(models) {
		take = len(models) - skip
	}

	modelList := ModelList{
		TotalPages: int(math.Ceil(float64(len(models)) / float64(pageSize))),
		Models:     models[skip : skip+take],
	}

	return modelList
}
