package main

import (
	"buggy/api/dashboard"
	"buggy/api/makes"
	"buggy/api/requestcontext"
	"buggy/internal/httpresponses"
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

	switch {
	case strings.HasPrefix(request.PathParameters["thepath"], "dashboard"):
		return dashboard.Handler(context)
	case strings.HasPrefix(request.PathParameters["thepath"], "makes"):
		return makes.Handler(context)
	}

	return httpresponses.NotFound, nil
}

func main() {
	lambda.Start(handler)
}
