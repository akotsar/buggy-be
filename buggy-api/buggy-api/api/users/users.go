package users

import (
	"buggy/api/requestcontext"
	"buggy/internal/auth"
	"buggy/internal/data/datacommon"
	"buggy/internal/data/userdata"
	"buggy/internal/httpresponses"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

type newUserRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type currentUserResponse struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	IsAdmin   bool   `json:"isAdmin"`
}

type profileResponse struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Gender    string `json:"gender"`
	Age       string `json:"age"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Hobby     string `json:"hobby"`
}

type updateProfileRequest struct {
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Gender          string `json:"gender"`
	Age             string `json:"age"`
	Address         string `json:"address"`
	Phone           string `json:"phone"`
	Hobby           string `json:"hobby"`
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

// Handler handles the dashboard-related API requests.
func Handler(context requestcontext.RequestContext) (events.APIGatewayProxyResponse, error) {
	switch {
	case strings.EqualFold(context.APIRequest.HTTPMethod, "POST") && context.Path == "users":
		return createUserHandler(context)
	case strings.EqualFold(context.APIRequest.HTTPMethod, "GET") && context.Path == "users/current":
		return getCurrentUserHandler(context)
	case strings.EqualFold(context.APIRequest.HTTPMethod, "GET") && context.Path == "users/profile":
		return getProfileHandler(context)
	case strings.EqualFold(context.APIRequest.HTTPMethod, "PUT") && context.Path == "users/profile":
		return updateProfileHandler(context)
	}

	return httpresponses.NotFound, nil
}

func createUserHandler(context requestcontext.RequestContext) (events.APIGatewayProxyResponse, error) {
	var request newUserRequest
	err := json.Unmarshal([]byte(context.APIRequest.Body), &request)
	if err != nil {
		log.Printf("Unable to deserialize request: %v\n", err)
		return httpresponses.InvalidRequest, nil
	}

	err = validateNewUserRequest(&request)
	if err != nil {
		log.Printf("Invalid request: %v\n", err)
		return httpresponses.CreateErrorResponse(400, err.Error()), nil
	}

	userID, err := auth.RegisterUser(context.Session, auth.RegisterUserInput{
		Username:  request.Username,
		Password:  request.Password,
		FirstName: request.FirstName,
		LastName:  request.LastName,
	})
	if err != nil {
		log.Printf("Unable to register the user: %v\n", err)
		return httpresponses.CreateErrorResponse(400, err.Error()), nil
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
		Username:  request.Username,
		IsAdmin:   false,
	}
	err = userdata.PutUser(context.Session, &userRecord)
	if err != nil {
		log.Printf("Unable to store user data: %v\n", err)
		return events.APIGatewayProxyResponse{}, err
	}

	return httpresponses.CreateJSONResponse(201, struct{}{}), nil
}

func getCurrentUserHandler(context requestcontext.RequestContext) (events.APIGatewayProxyResponse, error) {
	if len(context.UserID) == 0 {
		return httpresponses.Unauthorized, nil
	}

	userRecord, err := userdata.GetUserByID(context.Session, context.UserID)
	if err != nil {
		log.Printf("Unable to retrieve user information: %v\n", err)
		return events.APIGatewayProxyResponse{}, err
	}

	if userRecord == nil {
		log.Fatalln("Unable to retrieve user information: no user")
		return httpresponses.Unauthorized, nil
	}

	response := currentUserResponse{
		FirstName: userRecord.FirstName,
		LastName:  userRecord.LastName,
		IsAdmin:   userRecord.IsAdmin,
	}

	return httpresponses.CreateJSONResponse(200, &response), nil
}

func getProfileHandler(context requestcontext.RequestContext) (events.APIGatewayProxyResponse, error) {
	if len(context.UserID) == 0 {
		return httpresponses.Unauthorized, nil
	}

	userRecord, err := userdata.GetUserByID(context.Session, context.UserID)
	if err != nil {
		log.Printf("Unable to retrieve user information: %v\n", err)
		return events.APIGatewayProxyResponse{}, err
	}

	if userRecord == nil {
		log.Printf("Unable to retrieve user information: no user")
		return httpresponses.Unauthorized, nil
	}

	age := ""
	if userRecord.Age > 0 {
		age = fmt.Sprintf("%d", userRecord.Age)
	}
	response := profileResponse{
		Username:  userRecord.Username,
		FirstName: userRecord.FirstName,
		LastName:  userRecord.LastName,
		Gender:    userRecord.Gender,
		Age:       age,
		Address:   userRecord.Address,
		Phone:     userRecord.Phone,
		Hobby:     userRecord.Hobby,
	}

	return httpresponses.CreateJSONResponse(200, &response), nil
}

func updateProfileHandler(context requestcontext.RequestContext) (events.APIGatewayProxyResponse, error) {
	if len(context.UserID) == 0 {
		return httpresponses.Unauthorized, nil
	}

	var request updateProfileRequest
	err := json.Unmarshal([]byte(context.APIRequest.Body), &request)
	if err != nil {
		log.Printf("Unable to deserialize request: %v\n", err)
		return httpresponses.InvalidRequest, nil
	}

	err = validateUpdateProfileRequest(&request)
	if err != nil {
		log.Printf("Invalid request: %v\n", err)
		return httpresponses.CreateErrorResponse(400, err.Error()), nil
	}

	// Injecting bugs
	match, err := regexp.Match("[^A-Za-z0-9]", []byte(request.Age))
	if match {
		return httpresponses.CreateErrorResponse(400, "Get a candy ;)"), nil
	}

	if request.Hobby == "Knitting" {
		return httpresponses.CreateErrorResponse(400, "Knitting cannot be a hobby!"), nil
	}

	if len(request.Gender) > 10 {
		return httpresponses.CreateErrorResponse(400, "That's one weird gender!"), nil
	}

	if len(request.CurrentPassword) > 0 || len(request.NewPassword) > 0 {
		// Changing password.
		err = auth.ChangePassword(context.Session, &auth.ChangePasswordInput{
			Username:        context.Username,
			Token:           context.Token,
			CurrentPassword: request.CurrentPassword,
			NewPassword:     request.NewPassword,
		})
		if err != nil {
			return httpresponses.CreateErrorResponse(400, err.Error()), nil
		}
	}

	// Updating the profile
	userRecord, err := userdata.GetUserByID(context.Session, context.UserID)
	if err != nil {
		log.Printf("Unable to retrieve user: %v\n", err)
		return httpresponses.Unauthorized, nil
	}

	userRecord.FirstName = request.FirstName
	userRecord.LastName = request.LastName
	userRecord.Gender = request.Gender
	// Intentional bug
	age, err := strconv.Atoi(request.Age)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	userRecord.Age = age
	userRecord.Address = request.Address
	userRecord.Phone = request.Phone
	userRecord.Hobby = request.Hobby

	err = userdata.PutUser(context.Session, userRecord)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return httpresponses.CreateJSONResponse(200, struct{}{}), nil
}
