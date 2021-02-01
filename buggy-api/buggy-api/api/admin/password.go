package admin

import (
	"buggy/api/requestcontext"
	"buggy/internal/auth"
	"buggy/internal/httpresponses"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func changePasswordHandler(context requestcontext.RequestContext, username string) (events.APIGatewayProxyResponse, error) {
	var password string
	err := json.Unmarshal([]byte(context.APIRequest.Body), &password)
	if err != nil {
		return httpresponses.InvalidRequest, nil
	}

	err = auth.ChangePassword(context.Session, username, password)
	if err != nil {
		return httpresponses.CreateErrorResponse(400, err.Error()), nil
	}

	return httpresponses.CreateJSONResponse(200, struct{}{}), err
}
