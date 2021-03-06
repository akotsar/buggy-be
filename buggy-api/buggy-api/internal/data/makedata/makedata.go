package makedata

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

const makeType = "Make"

// MakeRecord represents a Make record in DynamoDB.
type MakeRecord struct {
	datacommon.DynamoRecordKey
	ShardID     uint8
	EntityType  string
	Name        string
	Image       string
	Description string
	Votes       int
}

// GetMakeByID returns make by its Id.
func GetMakeByID(session *session.Session, makeID string) (*MakeRecord, error) {
	dynamo := dynamodb.New(session)
	recordID := GenerateMakeRecordID(makeID)

	var response *MakeRecord = &MakeRecord{}
	exists, err := datacommon.GetItemByKey(dynamo, &datacommon.DynamoRecordKey{RecordID: recordID, TypeAndID: recordID}, response)
	if err != nil {
		log.Printf("Error while fetching user by id: %v", err)
		return nil, err
	}

	if !exists {
		response = nil
	}

	return response, nil
}

// GetTopMake returns a make with the most votes.
func GetTopMake(session *session.Session) (*MakeRecord, error) {
	dynamo := dynamodb.New(session)
	tableName := datacommon.TableName

	// Looking for the top make
	keyCondition := expression.Key("EntityType").Equal(expression.Value(makeType))

	expr, err := expression.NewBuilder().WithKeyCondition(keyCondition).Build()
	if err != nil {
		return nil, err
	}

	topMakeResult, err := dynamo.Query(&dynamodb.QueryInput{
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

	if len(topMakeResult.Items) <= 0 {
		// No makes are available
		log.Println("No makes found in the database.")
		return nil, nil
	}

	var make MakeRecord
	err = dynamodbattribute.UnmarshalMap(topMakeResult.Items[0], &make)
	if err != nil {
		return nil, err
	}

	return &make, nil
}

// GetMakesByIDs returns a list of makes by IDs.
func GetMakesByIDs(session *session.Session, makeIDs []string) ([]*MakeRecord, error) {
	dynamo := dynamodb.New(session)

	var makeRecords []*MakeRecord
	var recordIDs []string
	for _, id := range makeIDs {
		recordIDs = append(recordIDs, GenerateMakeRecordID(id))
	}

	err := datacommon.GetItemsByIDs(dynamo, recordIDs, &makeRecords)
	if err != nil {
		return nil, err
	}

	return makeRecords, nil
}

// PutMake writes a make record.
func PutMake(session *session.Session, make *MakeRecord) error {
	dynamo := dynamodb.New(session)

	make.EntityType = makeType

	_, err := datacommon.PutItem(dynamo, make)

	return err
}

// DeleteAllMakes deletes all makes form the database.
func DeleteAllMakes(session *session.Session) error {
	dynamo := dynamodb.New(session)

	return datacommon.DeleteAllByPrefix(dynamo, GenerateMakeRecordID(""))
}

// IncMakeVotes increments the number of votes for a make by one.
func IncMakeVotes(session *session.Session, makeID string) error {
	dynamo := dynamodb.New(session)

	recordID := GenerateMakeRecordID(makeID)
	return datacommon.IncField(dynamo, "Votes", recordID)
}

func GenerateMakeRecordID(ID string) string {
	return fmt.Sprintf("%s|%s", makeType, ID)
}
