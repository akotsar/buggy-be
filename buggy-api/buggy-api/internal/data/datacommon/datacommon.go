package datacommon

import (
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/rs/xid"
)

// DynamoRecordKey represents a basic record in Dynamo.
type DynamoRecordKey struct {
	RecordID  string
	TypeAndID string
}

// MaxShards defines maximum number of shards in Dynamo.
const MaxShards = 5

// TypeAndIDIndexName defines name of the type-and-id index.
const TypeAndIDIndexName = "TypeAndID"

// VotesIndexName defines name of the vodes index.
const VotesIndexName = "Votes"

// GetTableName returns the name of the Dynamo table.
func GetTableName() string {
	return os.Getenv("DATA_TABLE_NAME")
}

// GenerateNewShardID generates a new random shard ID.
func GenerateNewShardID() uint8 {
	return uint8(rand.Intn(MaxShards))
}

// GenerateNewID generates a new unique ID.
func GenerateNewID() string {
	return xid.New().String()
}

// GetIDFromTypeAndID returns ID of the record from TypeAndID value
func (record DynamoRecordKey) GetIDFromTypeAndID() string {
	return strings.SplitN(record.TypeAndID, "|", 2)[1]
}

// PutItem marshals and puts a recored into Dynamo
func PutItem(dynamo *dynamodb.DynamoDB, item interface{}) (*dynamodb.PutItemOutput, error) {
	itemmap, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return nil, err
	}

	return dynamo.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(GetTableName()),
		Item:      itemmap,
	})
}

// GetItemByKey reads a Dynamo item by its key.
func GetItemByKey(dynamo *dynamodb.DynamoDB, key *DynamoRecordKey, output interface{}) (bool, error) {
	keyMap, err := dynamodbattribute.MarshalMap(key)
	if err != nil {
		log.Fatalf("Unable to convert the key into an attribute map: %v", key)
		return false, err
	}

	request := &dynamodb.GetItemInput{TableName: aws.String(GetTableName()), Key: keyMap}
	result, err := dynamo.GetItem(request)
	if err != nil {
		log.Fatalf("Error while fetching record by id: %v", err)
		return false, err
	}

	var exists = false
	if result.Item != nil {
		exists = true
		dynamodbattribute.UnmarshalMap(result.Item, output)
	}

	return exists, nil
}

// GetItemsByKeyPrefix Finds Records with TypeAndID starting with a given prefix.
func GetItemsByKeyPrefix(dynamo *dynamodb.DynamoDB, prefix string, output interface{}) error {
	tableName := GetTableName()

	var shardOutputs [MaxShards]chan []map[string]*dynamodb.AttributeValue
	for i := range shardOutputs {
		ch := make(chan []map[string]*dynamodb.AttributeValue)
		shardOutputs[i] = ch

		keyCondition := expression.Key("ShardID").Equal(expression.Value(i)).And(
			expression.Key("TypeAndID").BeginsWith(prefix),
		)

		expr, err := expression.NewBuilder().WithKeyCondition(keyCondition).Build()
		if err != nil {
			log.Fatalf("Unable to generate key condition: %v\n", err)
			return err
		}

		go func() {
			defer func() { ch <- nil }()
			queryResult, err := dynamo.Query(&dynamodb.QueryInput{
				TableName:                 aws.String(tableName),
				IndexName:                 aws.String(TypeAndIDIndexName),
				KeyConditionExpression:    expr.KeyCondition(),
				ExpressionAttributeNames:  expr.Names(),
				ExpressionAttributeValues: expr.Values(),
			})

			if err != nil {
				log.Fatalf("Unable to generate key condition: %v\n", err)
				return
			}

			ch <- queryResult.Items
		}()
	}

	var allResults []map[string]*dynamodb.AttributeValue
	for _, shardResults := range shardOutputs {
		allResults = append(allResults, <-shardResults...)
	}

	err := dynamodbattribute.UnmarshalListOfMaps(allResults, output)
	if err != nil {
		log.Fatalf("Unable to unmarshall records: %v\n", err)
		return err
	}

	return nil
}
