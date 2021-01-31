package modeldata

import (
	"buggy/internal/data/datacommon"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

const modelType = "Model"

// ModelRecord represents a Model record in DynamoDB.
type ModelRecord struct {
	datacommon.DynamoRecordKey
	ShardID     uint8
	EntityType  string
	Name        string
	Image       string
	Description string
	EngineVol   float64
	MaxSpeed    int
	Votes       int
}

// GetMakeID returns make ID from TypeAndID.
func (model ModelRecord) GetMakeID() string {
	return strings.Split(model.TypeAndID, "|")[1]
}

// SortByVotesDescending implements sort.Interface for []VoteRecord based on the DateVoted field.
type SortByVotesDescending []ModelRecord

func (a SortByVotesDescending) Len() int           { return len(a) }
func (a SortByVotesDescending) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByVotesDescending) Less(i, j int) bool { return a[j].Votes < a[i].Votes }

// GetModelsByMakeID returns a list of models of a given make.
func GetModelsByMakeID(session *session.Session, makeID string) ([]ModelRecord, error) {
	dynamo := dynamodb.New(session)
	var models []ModelRecord

	err := datacommon.GetItemsByKeyPrefix(dynamo, GenerateModelRecordID(makeID), &models)
	if err != nil {
		return nil, err
	}

	return models, nil
}

// GetTopModel returns a model with the most votes.
func GetTopModel(session *session.Session) (*ModelRecord, error) {
	dynamo := dynamodb.New(session)
	tableName := datacommon.GetTableName()

	keyCondition := expression.Key("EntityType").Equal(expression.Value(modelType))

	expr, err := expression.NewBuilder().WithKeyCondition(keyCondition).Build()
	if err != nil {
		return nil, err
	}

	topModelResult, err := dynamo.Query(&dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		IndexName:                 aws.String(datacommon.VotesIndexName),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ScanIndexForward:          aws.Bool(false),
		Limit:                     aws.Int64(1),
	})
	if err != nil {
		return nil, err
	}

	if len(topModelResult.Items) <= 0 {
		// No makes are available
		log.Println("No makes found in the database.")
		return nil, nil
	}

	var make ModelRecord
	err = dynamodbattribute.UnmarshalMap(topModelResult.Items[0], &make)
	if err != nil {
		return nil, err
	}

	return &make, nil
}

// DeleteAllModels deletes all models form the database.
func DeleteAllModels(session *session.Session) error {
	dynamo := dynamodb.New(session)

	return datacommon.DeleteAllByPrefix(dynamo, GenerateModelRecordID(""))
}

// PutModels writes multiple model records.
func PutModels(session *session.Session, models []*ModelRecord) error {
	dynamo := dynamodb.New(session)

	modelAttrMaps := make([]map[string]*dynamodb.AttributeValue, 0, len(models))
	for _, m := range models {
		m.EntityType = modelType
		av, err := dynamodbattribute.MarshalMap(m)
		if err != nil {
			return err
		}

		modelAttrMaps = append(modelAttrMaps, av)
	}

	err := datacommon.PutItems(dynamo, modelAttrMaps)

	return err
}

// GenerateModelRecordID generates a Dynamo record ID for a given model.
func GenerateModelRecordID(modelID string) string {
	return fmt.Sprintf("%s|%s", modelType, modelID)
}
