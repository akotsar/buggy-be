package admin

import (
	"buggy/api/requestcontext"
	"buggy/internal/auth"
	"buggy/internal/httpresponses"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

func changePasswordHandler(context requestcontext.RequestContext, username string) (events.APIGatewayProxyResponse, error) {
	var password string
	err := json.Unmarshal([]byte(context.APIRequest.Body), &password)
	if err != nil {
		log.Printf("An error occurred while unmarshalling json: %v\n", err)
		return httpresponses.InvalidRequest, nil
	}

	err = auth.ChangePassword(context.Session, username, password)
	if err != nil {
		log.Printf("An error occurred while changing password: %v\n", err)
		return httpresponses.CreateErrorResponse(400, err.Error()), nil
	}

	return httpresponses.CreateJSONResponse(200, struct{}{}), err
}
