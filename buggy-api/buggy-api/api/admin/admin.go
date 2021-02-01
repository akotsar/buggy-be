package admin

import (
	"buggy/api/requestcontext"
	"buggy/internal/data/userdata"
	"buggy/internal/httpresponses"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

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

	log.Println(user)

	if !user.IsAdmin {
		return httpresponses.Unauthorized, nil
	}

	switch {
	case strings.EqualFold(context.APIRequest.HTTPMethod, "POST") && context.Path == "admin/reset/cars":
		return resetCarsHandler(context)
	}

	log.Println("No /admin route matched.")

	return httpresponses.NotFound, nil
}
