package main

import (
	"buggy/api/admin"
	"buggy/api/dashboard"
	"buggy/api/makes"
	"buggy/api/requestcontext"
	"buggy/internal/httpresponses"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	context := requestcontext.RequestContext{
		UserID:     "",
		APIRequest: &request,
		Path:       request.PathParameters["thepath"],
		Session:    session.Must(session.NewSession()),
	}

	log.Printf("Request: %s %s, UserId: %s, Body: %d", request.HTTPMethod, context.Path, context.UserID, len(request.Body))

	switch {
	case strings.HasPrefix(request.PathParameters["thepath"], "dashboard"):
		return dashboard.Handler(context)
	case strings.HasPrefix(request.PathParameters["thepath"], "makes"):
		return makes.Handler(context)
	case strings.HasPrefix(request.PathParameters["thepath"], "admin"):
		return admin.Handler(context)
	}

	return httpresponses.NotFound, nil
}

func main() {
	lambda.Start(handler)
}
