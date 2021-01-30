package main

import (
	"buggy/api/dashboard"
	"buggy/internal/httpresponses"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch {
	case strings.HasPrefix(request.PathParameters["thepath"], "dashboard"):
		return dashboard.Handler("", request)
	}

	return httpresponses.NotFound, nil
}

func main() {
	lambda.Start(handler)
}
