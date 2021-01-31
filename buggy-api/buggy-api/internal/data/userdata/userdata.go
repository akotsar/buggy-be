package userdata

import (
	"buggy/internal/data/datacommon"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const userType = "User"

// UserRecord represents a user.
type UserRecord struct {
	datacommon.DynamoRecordKey
	Username  string
	FirstName string
	LastName  string
	Gender    string
	Age       int
	Address   string
	Phone     string
	Hobby     string
	IsAdmin   bool
}

// PutUser writes a user record.
func PutUser(session *session.Session, user *UserRecord) error {
	dynamo := dynamodb.New(session)

	_, err := datacommon.PutItem(dynamo, user)

	return err
}

// GetUserByID returns a user by its ID.
func GetUserByID(session *session.Session, userID string) (*UserRecord, error) {
	dynamo := dynamodb.New(session)

	var userRecord UserRecord
	recordID := GenerateUserRecordID(userID)
	found, err := datacommon.GetItemByKey(dynamo, &datacommon.DynamoRecordKey{
		RecordID:  recordID,
		TypeAndID: recordID,
	}, &userRecord)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, nil
	}

	return &userRecord, nil
}

// GenerateUserRecordID creates new record ID for a user.
func GenerateUserRecordID(userID string) string {
	return fmt.Sprintf("%s|%s", userType, userID)
}
