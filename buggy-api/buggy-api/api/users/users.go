package users

import (
	"buggy/api/requestcontext"
	"buggy/internal/auth"
	"buggy/internal/data/datacommon"
	"buggy/internal/data/userdata"
	"buggy/internal/httpresponses"
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type newUserRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Handler handles the dashboard-related API requests.
func Handler(context requestcontext.RequestContext) (events.APIGatewayProxyResponse, error) {
	switch {
	case strings.EqualFold(context.APIRequest.HTTPMethod, "POST") && context.Path == "users":
		return createUserHandler(context)
	}

	return httpresponses.NotFound, nil
}

func createUserHandler(context requestcontext.RequestContext) (events.APIGatewayProxyResponse, error) {
	var request newUserRequest
	err := json.Unmarshal([]byte(context.APIRequest.Body), &request)
	if err != nil {
		log.Fatalf("Unable to deserialize request: %v\n", err)
		return httpresponses.InvalidRequest, nil
	}

	err = validateNewUserRequest(&request)
	if err != nil {
		log.Printf("Invalid request: %v\n", err)
		return httpresponses.CreateAPIResponse(400, []byte(err.Error())), nil
	}

	userID, err := auth.RegisterUser(context.Session, auth.RegisterUserInput{
		Username:  request.Username,
		Password:  request.Password,
		FirstName: request.FirstName,
		LastName:  request.LastName,
	})
	if err != nil {
		_, ok := err.(*cognitoidentityprovider.UsernameExistsException)
		if ok {
			log.Println(err)
			return httpresponses.CreateAPIResponse(400, []byte(err.Error())), nil
		}

		log.Fatalf("Unable to register the user: %v\n", err)
		return events.APIGatewayProxyResponse{}, err
	}

	// Registering the user in Dynamo
	recordID := userdata.GenerateUserRecordID(userID)
	userRecord := userdata.UserRecord{
		DynamoRecordKey: datacommon.DynamoRecordKey{
			RecordID:  recordID,
			TypeAndID: recordID,
		},
		FirstName: request.FirstName,
		LastName:  request.LastName,
	}
	err = userdata.PutUser(context.Session, &userRecord)
	if err != nil {
		log.Fatalf("Unable to store user data: %v\n", err)
		return events.APIGatewayProxyResponse{}, err
	}

	return httpresponses.CreateJSONResponse(201, struct{}{}), nil
}

func validateNewUserRequest(request *newUserRequest) error {
	if len(request.Username) <= 0 {
		return errors.New("username is required")
	}

	if len(request.Username) > 50 {
		return errors.New("username is too long")
	}

	if len(request.Password) <= 0 {
		return errors.New("password is required")
	}

	if len(request.Password) > 50 {
		return errors.New("password is too long")
	}

	if len(request.FirstName) <= 0 {
		return errors.New("first name is required")
	}

	if len(request.FirstName) > 250 {
		return errors.New("first name is too long")
	}

	if len(request.LastName) <= 0 {
		return errors.New("last name is required")
	}

	if len(request.LastName) > 250 {
		return errors.New("last name is too long")
	}

	return nil
}
