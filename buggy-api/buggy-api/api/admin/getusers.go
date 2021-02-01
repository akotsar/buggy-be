package admin

import (
	"buggy/api/requestcontext"
	"buggy/internal/auth"
	"buggy/internal/data/userdata"
	"buggy/internal/httpresponses"
	"log"
	"math/rand"
	"sort"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

var plainPasswords []string = []string{"123456", "password", "12345678", "qwerty", "123456789", "baseball", "dragon", "football", "monkey",
	"letmein", "abc123", "111111", "mustang", "access", "shadow", "master", "michael", "superman",
	"696969", "123123", "batman", "trustno1"}

type userItem struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	LockedOut bool   `json:"lockedOut"`
	Password  string `json:"password"`
	CanDelete bool   `json:"canDelete"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(plainPasswords), func(i, j int) { plainPasswords[i], plainPasswords[j] = plainPasswords[j], plainPasswords[i] })
}

func getUsersHandler(context requestcontext.RequestContext) (events.APIGatewayProxyResponse, error) {
	users, err := auth.GetAllUsers(context.Session)
	if err != nil {
		log.Printf("Unable to retrieve users: %v\n", err)
		return events.APIGatewayProxyResponse{}, err
	}

	var userIDs []string
	for _, u := range users {
		userIDs = append(userIDs, u.UserID)
	}

	userRecords, err := userdata.GetUsersByIDs(context.Session, userIDs)
	if err != nil {
		log.Printf("Unable to retrieve users: %v\n", err)
		return events.APIGatewayProxyResponse{}, err
	}

	var responses []userItem
	for i, u := range users {
		for _, ur := range userRecords {
			if u.UserID == ur.GetIDFromTypeAndID() {
				responses = append(responses, userItem{
					Username:  u.Username,
					FirstName: ur.FirstName,
					LastName:  ur.LastName,
					LockedOut: !u.Enabled,
					Password:  plainPasswords[i%len(plainPasswords)],
					CanDelete: u.UserID != context.UserID,
				})
				break
			}
		}
	}

	sort.Slice(responses, func(i, j int) bool {
		return responses[i].Username < responses[j].Username
	})

	return httpresponses.CreateJSONResponse(200, &responses), nil
}
