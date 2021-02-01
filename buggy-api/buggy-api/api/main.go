package main

import (
	"buggy/api/admin"
	"buggy/api/dashboard"
	"buggy/api/makes"
	"buggy/api/models"
	"buggy/api/oauth"
	"buggy/api/requestcontext"
	"buggy/api/users"
	"buggy/internal/auth"
	"buggy/internal/httpresponses"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
)

var authHeaderRegex = regexp.MustCompile("^Bearer (\\S+)$")

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	path, err := url.PathUnescape(request.PathParameters["thepath"])
	if err != nil {
		log.Printf("Unable to unescape the path: %v\n", err)
		path = request.PathParameters["thepath"]
	}

	context := requestcontext.RequestContext{
		APIRequest: &request,
		Path:       path,
		Session:    session.Must(session.NewSession()),
	}

	handleAuth(&context)

	log.Printf("Request: %s %s, UserId: %s, Body: %d", request.HTTPMethod, context.Path, context.UserID, len(request.Body))

	switch {
	case strings.HasPrefix(context.Path, "dashboard"):
		return dashboard.Handler(context)
	case strings.HasPrefix(context.Path, "makes"):
		return makes.Handler(context)
	case strings.HasPrefix(context.Path, "models"):
		return models.Handler(context)
	case strings.HasPrefix(context.Path, "admin"):
		return admin.Handler(context)
	case strings.HasPrefix(context.Path, "users"):
		return users.Handler(context)
	case strings.HasPrefix(context.Path, "oauth"):
		return oauth.Handler(context)
	}

	return httpresponses.NotFound, nil
}

func handleAuth(context *requestcontext.RequestContext) {
	authHeader, ok := context.APIRequest.Headers["Authorization"]
	if !ok {
		// The user is not authenticated.
		return
	}

	match := authHeaderRegex.FindStringSubmatch(authHeader)
	if match == nil {
		log.Printf("Invalid Authorization header: %s\n", authHeader)
		return
	}

	token := match[1]
	validationOutput, err := auth.ValidateToken(token)
	if err != nil {
		log.Printf("Token validation failed: %v\n", err)
		return
	}

	context.UserID = validationOutput.UserID
	context.Username = validationOutput.Username
	context.Token = validationOutput.Token
}

func main() {
	lambda.Start(handler)
}
