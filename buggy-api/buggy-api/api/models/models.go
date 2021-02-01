package models

import (
	"buggy/api/requestcontext"
	"buggy/internal/httpresponses"
	"log"
	"regexp"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

var getModelByIDRegexp *regexp.Regexp = regexp.MustCompile("^models\\/([^\\/]+)$")
var voteModelRegexp *regexp.Regexp = regexp.MustCompile("^models\\/([^\\/]+)\\/vote$")

// Handler handles the user-related API requests.
func Handler(context requestcontext.RequestContext) (events.APIGatewayProxyResponse, error) {
	getModelByIDMatch := getModelByIDRegexp.FindStringSubmatch(context.Path)
	voteModelMatch := voteModelRegexp.FindStringSubmatch(context.Path)

	switch {
	case strings.EqualFold(context.APIRequest.HTTPMethod, "GET") && context.Path == "models":
		return getModelsHandler(context)
	case strings.EqualFold(context.APIRequest.HTTPMethod, "GET") && getModelByIDMatch != nil:
		return getModelByIDHandler(context, getModelByIDMatch[1])
	case strings.EqualFold(context.APIRequest.HTTPMethod, "POST") && voteModelMatch != nil:
		return voteHandler(context, voteModelMatch[1])
	}

	log.Println("No /models route matched.")

	return httpresponses.NotFound, nil
}
