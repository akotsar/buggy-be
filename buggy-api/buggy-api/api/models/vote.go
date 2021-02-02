package models

import (
	"buggy/api/requestcontext"
	"buggy/internal/data/datacommon"
	"buggy/internal/data/makedata"
	"buggy/internal/data/modeldata"
	"buggy/internal/data/votedata"
	"buggy/internal/httpresponses"
	"encoding/json"
	"errors"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

type voteRequest struct {
	Comment string `json:"comment"`
}

func voteHandler(context requestcontext.RequestContext, modelID string) (events.APIGatewayProxyResponse, error) {
	if len(context.UserID) == 0 {
		return httpresponses.Unauthorized, nil
	}

	var request voteRequest
	err := json.Unmarshal([]byte(context.APIRequest.Body), &request)
	if err != nil {
		log.Printf("Unable to deserialize request: %v\n", err)
		return httpresponses.InvalidRequest, nil
	}

	err = validateVoteRequest(&request)
	if err != nil {
		log.Printf("Invalid request: %v\n", err)
		return httpresponses.CreateErrorResponse(400, err.Error()), nil
	}

	voted, err := votedata.HasUserVotedForModel(context.Session, modelID, context.UserID)
	if err != nil {
		log.Printf("Unable to determine if the user has voted: %v\n", err)
		return events.APIGatewayProxyResponse{}, err
	}

	if voted {
		log.Println("Cannot vote more than once.")
		return httpresponses.CreateErrorResponse(400, "Cannot vote more than once"), nil
	}

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		modelRecordID := modeldata.GenerateModelRecordID(modelID)
		voteRecordID := votedata.GenerateVoteRecordID(modelID, context.UserID)
		vote := votedata.VoteRecord{
			DynamoRecordKey: datacommon.DynamoRecordKey{RecordID: modelRecordID, TypeAndID: voteRecordID},
			Comment:         request.Comment,
			DateVoted:       time.Now(),
		}
		err = votedata.PutVote(context.Session, &vote)
		if err != nil {
			log.Printf("Unable to save a vote: %v\n", err)
		}
	}()

	go func() {
		defer wg.Done()
		err := modeldata.IncModelVotes(context.Session, modelID)
		if err != nil {
			log.Printf("Unable to increment model votes: %v\n", err)
		}
	}()

	go func() {
		defer wg.Done()
		makeID := strings.Split(modelID, "|")[0]
		err := makedata.IncMakeVotes(context.Session, makeID)
		if err != nil {
			log.Printf("Unable to increment make votes: %v\n", err)
		}
	}()

	wg.Wait()

	return httpresponses.CreateJSONResponse(200, struct{}{}), nil
}

func validateVoteRequest(request *voteRequest) error {
	if len(request.Comment) > 500 {
		return errors.New("comment is too long")
	}

	return nil
}
