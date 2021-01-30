package admin

import (
	"buggy/api/requestcontext"
	"buggy/internal/httpresponses"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

// Handler handles the admin-related API requests.
func Handler(context requestcontext.RequestContext) (events.APIGatewayProxyResponse, error) {
	switch {
	case strings.EqualFold(context.APIRequest.HTTPMethod, "POST") && context.Path == "admin/reset/cars":
		return resetCarsHandler(context)
	}

	return httpresponses.NotFound, nil
}
