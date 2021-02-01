package makes

import (
	"buggy/api/dtos"
	"buggy/api/requestcontext"
	"buggy/internal/data/makedata"
	"buggy/internal/data/modeldata"
	"buggy/internal/httpresponses"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

const defaultPageSize = 5

var getMakeByIDRegexp *regexp.Regexp

func init() {
	getMakeByIDRegexp = regexp.MustCompile("^makes\\/(\\w+)$")
}

// Handler handles the user-related API requests.
func Handler(context requestcontext.RequestContext) (events.APIGatewayProxyResponse, error) {
	getMakeByIDMatch := getMakeByIDRegexp.FindStringSubmatch(context.Path)

	switch {
	case strings.EqualFold(context.APIRequest.HTTPMethod, "GET") && getMakeByIDMatch != nil:
		return getMakeByIDHandler(context, getMakeByIDMatch[1])
	}

	log.Println("No /makes route matched.")

	return httpresponses.NotFound, nil
}

func getMakeByIDHandler(context requestcontext.RequestContext, makeID string) (events.APIGatewayProxyResponse, error) {
	makeRecord, err := makedata.GetMakeByID(context.Session, makeID)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if makeRecord == nil {
		return httpresponses.NotFound, nil
	}

	modelRecords, err := modeldata.GetModelsByMakeID(context.Session, makeID)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	var models []*dtos.ModelItem
	for _, modelRecord := range modelRecords {
		model := dtos.NewModelItemFromRecord(modelRecord)
		model.Make = makeRecord.Name
		model.MakeImage = makeRecord.Image
		models = append(models, model)
	}

	response := dtos.NewMakeFromRecord(makeRecord)
	response.Models = getModelsPage(context, models)

	// Retrieving comments.
	dtos.GetCommentsForModels(context.Session, response.Models.Models)

	return httpresponses.CreateJSONResponse(200, response), nil
}

func getModelsPage(context requestcontext.RequestContext, models []*dtos.ModelItem) dtos.ModelList {
	// Pagination
	page := context.APIRequest.QueryStringParameters["modelsPage"]
	pageIndex, err := strconv.Atoi(page)
	if err != nil {
		pageIndex = 1
	}

	orderBy := context.APIRequest.QueryStringParameters["modelsOrderBy"]

	modelList := dtos.GetModelsPage(models, pageIndex, defaultPageSize, orderBy)

	return modelList
}
