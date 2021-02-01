package admin

import (
	"buggy/api/requestcontext"
	"buggy/internal/data/userdata"
	"buggy/internal/httpresponses"
	"log"
	"regexp"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

var lockUserRegexp *regexp.Regexp = regexp.MustCompile("^admin\\/users\\/([^\\/]+)\\/lock$")
var unlockUserRegexp *regexp.Regexp = regexp.MustCompile("^admin\\/users\\/([^\\/]+)\\/unlock$")
var deleteUserRegexp *regexp.Regexp = regexp.MustCompile("^admin\\/users\\/([^\\/]+)$")
var changePasswordRegexp *regexp.Regexp = regexp.MustCompile("^admin\\/users\\/([^\\/]+)\\/password$")

// Handler handles the admin-related API requests.
func Handler(context requestcontext.RequestContext) (events.APIGatewayProxyResponse, error) {
	if len(context.UserID) == 0 {
		return httpresponses.Unauthorized, nil
	}

	user, err := userdata.GetUserByID(context.Session, context.UserID)
	if err != nil {
		log.Printf("Unable to retrieve user by ID: %v\n", err)
		return httpresponses.Unauthorized, nil
	}

	if !user.IsAdmin {
		return httpresponses.Unauthorized, nil
	}

	lockUserMatch := lockUserRegexp.FindStringSubmatch(context.Path)
	unlockUserMatch := unlockUserRegexp.FindStringSubmatch(context.Path)
	deleteUserMatch := deleteUserRegexp.FindStringSubmatch(context.Path)
	changePasswordMatch := changePasswordRegexp.FindStringSubmatch(context.Path)

	switch {
	case strings.EqualFold(context.APIRequest.HTTPMethod, "POST") && context.Path == "admin/reset/cars":
		return resetCarsHandler(context)
	case strings.EqualFold(context.APIRequest.HTTPMethod, "GET") && context.Path == "admin/users":
		return getUsersHandler(context)
	case strings.EqualFold(context.APIRequest.HTTPMethod, "PUT") && lockUserMatch != nil:
		return lockUserHandler(context, lockUserMatch[1])
	case strings.EqualFold(context.APIRequest.HTTPMethod, "PUT") && unlockUserMatch != nil:
		return unlockUserHandler(context, unlockUserMatch[1])
	case strings.EqualFold(context.APIRequest.HTTPMethod, "PUT") && changePasswordMatch != nil:
		return changePasswordHandler(context, changePasswordMatch[1])
	case strings.EqualFold(context.APIRequest.HTTPMethod, "DELETE") && deleteUserMatch != nil:
		return deleteUserHandler(context, deleteUserMatch[1])
	}

	log.Println("No /admin route matched.")

	return httpresponses.NotFound, nil
}
