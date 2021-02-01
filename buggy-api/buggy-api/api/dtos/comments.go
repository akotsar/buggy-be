package dtos

import (
	"buggy/internal/data/votedata"
	"log"
	"sort"

	"github.com/aws/aws-sdk-go/aws/session"
)

const maxComments = 3

// GetCommentsForModels retrieves top X comments for each of the model in the list.
func GetCommentsForModels(session *session.Session, models []*ModelItem) {
	voteChannels := make([]chan []*votedata.VoteRecord, len(models))
	for i, model := range models {
		modelID := model.ID
		voteChannels[i] = make(chan []*votedata.VoteRecord)
		ch := voteChannels[i]
		go func() {
			defer func() { ch <- nil }()
			result, err := votedata.GetNonEmptyCommentsByModelID(session, modelID)
			if err != nil {
				log.Printf("Unable to fetch votes: %v\n", err)
				return
			}

			ch <- result
		}()
	}

	for i, ch := range voteChannels {
		comments := <-ch

		models[i].TotalComments = len(comments)
		models[i].Comments = make([]string, 0)

		sort.Sort(votedata.SortByDateDescending(comments))
		if len(comments) > maxComments {
			comments = comments[:maxComments]
		}
		for _, c := range comments {
			models[i].Comments = append(models[i].Comments, c.Comment)
		}
	}
}
