package votedata

import (
	"buggy/internal/data/datacommon"
	"buggy/internal/data/modeldata"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const modelType = "Vote"

// VoteRecord represents a user vote for a given car model.
type VoteRecord struct {
	datacommon.DynamoRecordKey
	Comment   string
	DateVoted time.Time
}

// SortByDateDescending implements sort.Interface for []VoteRecord based on the DateVoted field.
type SortByDateDescending []VoteRecord

func (a SortByDateDescending) Len() int           { return len(a) }
func (a SortByDateDescending) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByDateDescending) Less(i, j int) bool { return a[j].DateVoted.Before(a[i].DateVoted) }

// GetMakeID returns make ID from TypeAndID.
func (model VoteRecord) GetMakeID() string {
	return strings.Split(model.TypeAndID, "|")[1]
}

// GetModelID returns model ID from TypeAndID.
func (model VoteRecord) GetModelID() string {
	return strings.Split(model.TypeAndID, "|")[2]
}

// GetUserID returns user ID from TypeAndID.
func (model VoteRecord) GetUserID() string {
	return strings.Split(model.TypeAndID, "|")[3]
}

// GetVotesByModelID returns a list of user votes for a given car model.
func GetVotesByModelID(session *session.Session, modelID string) (*[]VoteRecord, error) {
	dynamo := dynamodb.New(session)
	var votes []VoteRecord

	err := datacommon.GetItemsByKeyPrefix(dynamo, generateVoteRecordID(modelID, ""), &votes)
	if err != nil {
		return nil, err
	}

	return &votes, nil
}

// HasUserVotedForModel returns true if the given user has voted for the given model.
func HasUserVotedForModel(session *session.Session, modelID string, userID string) (bool, error) {
	dynamo := dynamodb.New(session)
	var vote VoteRecord

	exists, err := datacommon.GetItemByKey(dynamo, &datacommon.DynamoRecordKey{
		RecordID:  modeldata.GenerateModelRecordID(modelID),
		TypeAndID: generateVoteRecordID(modelID, userID),
	}, &vote)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func generateVoteRecordID(modelID, userID string) string {
	return fmt.Sprintf("%s|%s|%s", modelType, modelID, userID)
}
