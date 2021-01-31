package oauth

import (
	"buggy/api/requestcontext"
	"buggy/internal/auth"
	"buggy/internal/httpresponses"
	"log"
	"net/url"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gorilla/schema"
)

type tokenRequest struct {
	GrantType string `schema:"grant_type"`
	Username  string `schema:"username,required"`
	Password  string `schema:"password,required"`
}

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

var decoder = schema.NewDecoder()

// Handler handles the dashboard-related API requests.
func Handler(context requestcontext.RequestContext) (events.APIGatewayProxyResponse, error) {
	switch {
	case strings.EqualFold(context.APIRequest.HTTPMethod, "POST") && context.Path == "oauth/token":
		return authenticateHandler(context)
	}

	return httpresponses.NotFound, nil
}

func authenticateHandler(context requestcontext.RequestContext) (events.APIGatewayProxyResponse, error) {
	parsedQuery, err := url.ParseQuery(context.APIRequest.Body)
	if err != nil {
		log.Fatalf("Invalid oauth query: %v\n", err)
		return httpresponses.InvalidRequest, nil
	}

	var request tokenRequest
	err = decoder.Decode(&request, parsedQuery)
	if err != nil {
		log.Fatalf("Invalid oauth query: %v\n", err)
		return httpresponses.InvalidRequest, nil
	}

	result, err := auth.LoginUser(context.Session, auth.LoginUserInput{
		Username: request.Username,
		Password: request.Password,
	})
	if err != nil {
		log.Printf("Unable to authenticate the user: %v\n", err)
		return httpresponses.CreateAPIResponse(401, []byte("Invalid credentials")), nil
	}

	response := tokenResponse{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
		TokenType:    result.TokenType,
		ExpiresIn:    result.ExpiresIn,
	}

	return httpresponses.CreateJSONResponse(200, response), nil
}
