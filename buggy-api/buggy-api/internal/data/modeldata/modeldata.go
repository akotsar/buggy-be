package modeldata

import (
	"buggy/internal/data/datacommon"
	"fmt"
	"log"

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
	ModelID     string
	Name        string
	Image       string
	Description string
	EngineVol   float64
	MaxSpeed    int
	Votes       int
}

// GetTopModel returns a model with the most votes.
func GetTopModel(session *session.Session) (*ModelRecord, error) {
	dynamo := dynamodb.New(session)
	tableName := datacommon.GetTableName()

	// Looking for the top make
	keyCondition := expression.Key("EntityType").Equal(expression.Value(modelType))

	expr, err := expression.NewBuilder().WithKeyCondition(keyCondition).Build()
	if err != nil {
		log.Fatalf("Unable to generate key condition: %v\n", err)
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
		log.Fatalf("Unable to query top model records: %v\n", err)
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
		log.Fatalf("Unable to unmarshall record: %v\n", err)
		return nil, err
	}

	return &make, nil
}

func generateModelRecord(ID string) string {
	return fmt.Sprintf("%s|%s", modelType, ID)
}
