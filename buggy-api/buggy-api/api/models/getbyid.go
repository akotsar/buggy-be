package models

import (
	"buggy/api/dtos"
	"buggy/api/requestcontext"
	"buggy/internal/data/makedata"
	"buggy/internal/data/modeldata"
	"buggy/internal/data/userdata"
	"buggy/internal/data/votedata"
	"buggy/internal/httpresponses"
	"fmt"
	"log"
	"sort"

	"github.com/aws/aws-lambda-go/events"
)

func getModelByIDHandler(context requestcontext.RequestContext, modelID string) (events.APIGatewayProxyResponse, error) {
	modelRecord, err := modeldata.GetModelByID(context.Session, modelID)
	if err != nil {
		log.Printf("Unable to retrieve model by id: %v\n", err)
		return events.APIGatewayProxyResponse{}, err
	}

	if modelRecord == nil {
		return httpresponses.NotFound, nil
	}

	response := dtos.NewModelDetailsFromRecord(modelRecord)
	makeRecord, err := makedata.GetMakeByID(context.Session, response.MakeID)
	if err != nil {
		log.Printf("Unable to retrieve make by id: %v\n", err)
		return events.APIGatewayProxyResponse{}, err
	}

	if makeRecord == nil {
		return httpresponses.NotFound, nil
	}

	response.Make = makeRecord.Name
	response.MakeImage = makeRecord.Image
	response.Comments = getModelComments(context, modelID)

	if len(context.UserID) > 0 {
		voted, err := votedata.HasUserVotedForModel(context.Session, modelID, context.UserID)
		if err != nil {
			log.Printf("Unable to determins if the user has voted: %v\n", err)
		}

		response.CanVote = !voted
	}

	return httpresponses.CreateJSONResponse(200, response), nil
}

func getModelComments(context requestcontext.RequestContext, modelID string) []*dtos.ModelComment {
	var comments []*dtos.ModelComment

	commentRecords, err := votedata.GetNonEmptyCommentsByModelID(context.Session, modelID)
	if err != nil {
		log.Printf("Unable to retrieve votes by model id: %v\n", err)
		return comments
	}

	seenUserIDs := make(map[string]bool)
	var userIDs []string
	sort.Sort(votedata.SortByDateDescending(commentRecords))
	for _, c := range commentRecords {
		userID := c.GetUserID()
		if !seenUserIDs[userID] {
			seenUserIDs[userID] = true
			userIDs = append(userIDs, userID)
		}
	}

	// Fetching user information.
	userRecords, err := userdata.GetUsersByIDs(context.Session, userIDs)
	if err != nil {
		log.Printf("Unable to retrieve users ids: %v\n", err)
	}

	userFullNamesByID := make(map[string]string)
	for _, r := range userRecords {
		userFullNamesByID[r.GetIDFromTypeAndID()] = fmt.Sprintf("%s %s", r.FirstName, r.LastName)
	}

	for _, c := range commentRecords {
		userID := c.GetUserID()
		comments = append(comments, &dtos.ModelComment{
			User:       userFullNamesByID[userID],
			DatePosted: c.DateVoted,
			Text:       c.Comment,
		})
	}

	return comments
}
