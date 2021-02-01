package votedata

import (
	"buggy/internal/data/datacommon"
	"buggy/internal/data/modeldata"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

const voteType = "Vote"

// VoteRecord represents a user vote for a given car model.
type VoteRecord struct {
	datacommon.DynamoRecordKey
	Comment   string
	DateVoted time.Time
}

// SortByDateDescending implements sort.Interface for []VoteRecord based on the DateVoted field.
type SortByDateDescending []*VoteRecord

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

// GetNonEmptyCommentsByModelID returns a list of non-empty user comments for a given car model.
func GetNonEmptyCommentsByModelID(session *session.Session, modelID string) ([]*VoteRecord, error) {
	dynamo := dynamodb.New(session)

	keyCondition := expression.Key("RecordID").Equal(expression.Value(modeldata.GenerateModelRecordID(modelID))).And(
		expression.Key("TypeAndID").BeginsWith(GenerateVoteRecordID(modelID, "")),
	)

	nonEmptyCondition := expression.Name("Comment").AttributeNotExists().Or(
		expression.Name("Comment").NotEqual(expression.Value("")))

	expr, err := expression.NewBuilder().WithKeyCondition(keyCondition).WithFilter(nonEmptyCondition).Build()
	if err != nil {
		return nil, err
	}

	input := &dynamodb.QueryInput{
		TableName:                 aws.String(datacommon.TableName),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}
	queryResult, err := dynamo.Query(input)
	if err != nil {
		return nil, err
	}

	var votes []*VoteRecord
	err = dynamodbattribute.UnmarshalListOfMaps(queryResult.Items, &votes)
	if err != nil {
		return nil, err
	}

	return votes, nil
}

// HasUserVotedForModel returns true if the given user has voted for the given model.
func HasUserVotedForModel(session *session.Session, modelID string, userID string) (bool, error) {
	dynamo := dynamodb.New(session)
	var vote VoteRecord

	exists, err := datacommon.GetItemByKey(dynamo, &datacommon.DynamoRecordKey{
		RecordID:  modeldata.GenerateModelRecordID(modelID),
		TypeAndID: GenerateVoteRecordID(modelID, userID),
	}, &vote)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// DeleteAllVotes deletes all votes form the database.
func DeleteAllVotes(session *session.Session) error {
	dynamo := dynamodb.New(session)

	return datacommon.DeleteAllByPrefix(dynamo, fmt.Sprintf("%s|", voteType))
}

// PutVotes writes multiple vote records.
func PutVotes(session *session.Session, votes []*VoteRecord) error {
	dynamo := dynamodb.New(session)

	attrMaps := make([]map[string]*dynamodb.AttributeValue, 0, len(votes))
	for _, v := range votes {
		av, err := dynamodbattribute.MarshalMap(v)
		if err != nil {
			return err
		}

		attrMaps = append(attrMaps, av)
	}

	err := datacommon.PutItems(dynamo, attrMaps)

	return err
}

// PutVote creates a new vote record.
func PutVote(session *session.Session, vote *VoteRecord) error {
	dynamo := dynamodb.New(session)
	_, err := datacommon.PutItem(dynamo, vote)
	return err
}

func GenerateVoteRecordID(modelID, userID string) string {
	return fmt.Sprintf("%s|%s|%s", voteType, modelID, userID)
}
