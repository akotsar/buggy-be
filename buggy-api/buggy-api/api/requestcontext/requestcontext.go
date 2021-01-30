package requestcontext

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/session"
)

// RequestContext contains context information of an API request.
type RequestContext struct {
	UserID     string
	Path       string
	APIRequest *events.APIGatewayProxyRequest
	Session    *session.Session
}
