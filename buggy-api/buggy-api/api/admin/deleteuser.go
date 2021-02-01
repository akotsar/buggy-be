package admin

import (
	"buggy/api/requestcontext"
	"buggy/internal/auth"
	"buggy/internal/data/userdata"
	"buggy/internal/httpresponses"

	"github.com/aws/aws-lambda-go/events"
)

func deleteUserHandler(context requestcontext.RequestContext, username string) (events.APIGatewayProxyResponse, error) {
	userInfo, err := auth.GetUserByUsername(context.Session, username)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if context.UserID == userInfo.UserID {
		// Cannot delete self.
		return httpresponses.InvalidRequest, nil
	}

	err = userdata.DeleteUser(context.Session, userInfo.UserID)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	err = auth.DeleteUser(context.Session, username)
	return httpresponses.CreateJSONResponse(200, struct{}{}), err
}
