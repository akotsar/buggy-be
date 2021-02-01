package admin

import (
	"buggy/api/requestcontext"
	"buggy/internal/auth"
	"buggy/internal/httpresponses"

	"github.com/aws/aws-lambda-go/events"
)

func lockUserHandler(context requestcontext.RequestContext, username string) (events.APIGatewayProxyResponse, error) {
	err := auth.LockUser(context.Session, username)
	return httpresponses.CreateJSONResponse(200, struct{}{}), err
}

func unlockUserHandler(context requestcontext.RequestContext, username string) (events.APIGatewayProxyResponse, error) {
	err := auth.UnlockUser(context.Session, username)
	return httpresponses.CreateJSONResponse(200, struct{}{}), err
}
