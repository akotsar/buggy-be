package makes

import (
	"buggy/api/dtos"
	"buggy/api/requestcontext"
	"buggy/internal/data/makedata"
	"buggy/internal/data/modeldata"
	"buggy/internal/data/votedata"
	"buggy/internal/httpresponses"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/session"
)

const defaultPageSize = 5

var getMakeByIDRegexp *regexp.Regexp

func init() {
	getMakeByIDRegexp = regexp.MustCompile("^makes\\/(\\w+)")
}

// Handler handles the user-related API requests.
func Handler(context requestcontext.RequestContext) (events.APIGatewayProxyResponse, error) {
	getMakeByIDMatch := getMakeByIDRegexp.FindStringSubmatch(context.Path)

	switch {
	case strings.EqualFold(context.APIRequest.HTTPMethod, "GET") && getMakeByIDMatch != nil:
		return getMakeByIDHandler(context, getMakeByIDMatch[1])
	}

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

	response := dtos.NewMakeFromRecord(makeRecord)
	response.Models = getModelsPage(context, *makeRecord, modelRecords)

	// Retrieving comments.
	getCommentsForModels(context.Session, response.Models.Models)

	return httpresponses.CreateJSONResponse(200, response), nil
}

func getModelsPage(context requestcontext.RequestContext, makeRecord makedata.MakeRecord, modelRecords []modeldata.ModelRecord) dtos.ModelList {
	var models []*dtos.ModelItem
	for _, modelRecord := range modelRecords {
		model := dtos.NewModelItemFromRecord(&modelRecord)
		model.Make = makeRecord.Name
		model.MakeImage = makeRecord.Image
		models = append(models, model)
	}

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

func getCommentsForModels(session *session.Session, models []*dtos.ModelItem) {
	voteChannels := make([]chan *[]votedata.VoteRecord, len(models))
	for i, model := range models {
		modelID := model.ID
		voteChannels[i] = make(chan *[]votedata.VoteRecord)
		ch := voteChannels[i]
		go func() {
			defer func() { ch <- nil }()
			result, err := votedata.GetVotesByModelID(session, modelID)
			if err != nil {
				log.Printf("Unable to fetch votes: %v\n", err)
				return
			}

			ch <- result
		}()
	}

	for i, ch := range voteChannels {
		comments := <-ch

		var nonEmptyComments []votedata.VoteRecord
		for _, c := range *comments {
			if len(c.Comment) > 0 {
				nonEmptyComments = append(nonEmptyComments, c)
			}
		}

		models[i].TotalComments = len(nonEmptyComments)
		models[i].Comments = make([]string, 0)

		sort.Sort(votedata.SortByDateDescending(nonEmptyComments))
		if len(nonEmptyComments) > 3 {
			nonEmptyComments = nonEmptyComments[:3]
		}
		for _, c := range nonEmptyComments {
			models[i].Comments = append(models[i].Comments, c.Comment)
		}
	}
}
