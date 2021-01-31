package datacommon

import (
	"log"
	"math"
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
		return false, err
	}

	request := &dynamodb.GetItemInput{TableName: aws.String(GetTableName()), Key: keyMap}
	result, err := dynamo.GetItem(request)
	if err != nil {
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
			return err
		}

		go func() {
			defer func() { ch <- nil }()
			input := &dynamodb.QueryInput{
				TableName:                 aws.String(tableName),
				IndexName:                 aws.String(TypeAndIDIndexName),
				KeyConditionExpression:    expr.KeyCondition(),
				ExpressionAttributeNames:  expr.Names(),
				ExpressionAttributeValues: expr.Values(),
			}
			queryResult, err := dynamo.Query(input)

			if err != nil {
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
		return err
	}

	return nil
}

// DeleteAllByPrefix deletes all items with a given prefix form the database.
func DeleteAllByPrefix(dynamo *dynamodb.DynamoDB, prefix string) error {
	var items []DynamoRecordKey
	err := GetItemsByKeyPrefix(dynamo, prefix, &items)
	if err != nil {
		return err
	}

	log.Printf("Deleting %d items.", len(items))
	writeRequests := make([]*dynamodb.WriteRequest, 0, len(items))
	for _, item := range items {
		log.Printf("Deleting %v.", item)

		key, _ := dynamodbattribute.MarshalMap(&item)
		r := dynamodb.WriteRequest{
			DeleteRequest: &dynamodb.DeleteRequest{
				Key: key,
			},
		}

		writeRequests = append(writeRequests, &r)
	}

	const maxWriteOps = 25

	batches := int(math.Ceil(float64(len(items)) / maxWriteOps))
	for i := 0; i < batches; i++ {
		start, end := i*maxWriteOps, (i+1)*maxWriteOps
		if end > len(writeRequests) {
			end = len(writeRequests)
		}
		ops := writeRequests[start:end]

		_, err = dynamo.BatchWriteItem(&dynamodb.BatchWriteItemInput{
			RequestItems: map[string][]*dynamodb.WriteRequest{
				GetTableName(): ops,
			},
		})
		if err != nil {
			log.Printf("Unable to delete records: %v\n", err)
		}
	}

	return nil
}

// PutItems writes multiple records into the database.
func PutItems(dynamo *dynamodb.DynamoDB, items []map[string]*dynamodb.AttributeValue) error {
	log.Printf("Writing %d items.", len(items))
	writeRequests := make([]*dynamodb.WriteRequest, 0, len(items))
	for _, item := range items {
		r := dynamodb.WriteRequest{
			PutRequest: &dynamodb.PutRequest{
				Item: item,
			},
		}

		writeRequests = append(writeRequests, &r)
	}

	const maxWriteOps = 25

	batches := int(math.Ceil(float64(len(items)) / maxWriteOps))
	for i := 0; i < batches; i++ {
		start, end := i*maxWriteOps, (i+1)*maxWriteOps
		if end > len(writeRequests) {
			end = len(writeRequests)
		}
		ops := writeRequests[start:end]

		_, err := dynamo.BatchWriteItem(&dynamodb.BatchWriteItemInput{
			RequestItems: map[string][]*dynamodb.WriteRequest{
				GetTableName(): ops,
			},
		})
		if err != nil {
			log.Printf("Unable to write records: %v\n", err)
		}
	}

	return nil
}
