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
	FirstName string
	LastName  string
	Gender    string
	Age       int
	Address   string
	Phone     string
	Hobby     string
}

// PutUser writes a user record.
func PutUser(session *session.Session, user *UserRecord) error {
	dynamo := dynamodb.New(session)

	_, err := datacommon.PutItem(dynamo, user)

	return err
}

// GenerateUserRecordID creates new record ID for a user.
func GenerateUserRecordID(userID string) string {
	return fmt.Sprintf("%s|%s", userType, userID)
}
