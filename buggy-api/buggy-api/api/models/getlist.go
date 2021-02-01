package models

import (
	"buggy/api/dtos"
	"buggy/api/requestcontext"
	"buggy/internal/data/makedata"
	"buggy/internal/data/modeldata"
	"buggy/internal/httpresponses"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

const defaultPageSize = 5

func getModelsHandler(context requestcontext.RequestContext) (events.APIGatewayProxyResponse, error) {
	var modelRecords []*modeldata.ModelRecord
	var err error
	makeIDParam := context.APIRequest.QueryStringParameters["makeId"]
	if len(makeIDParam) > 0 {
		modelRecords, err = modeldata.GetModelsByMakeID(context.Session, makeIDParam)
	} else {
		modelRecords, err = modeldata.GetAllModels(context.Session)
	}
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	seenMakeIDs := make(map[string]bool)
	var makeIDs []string
	var models []*dtos.ModelItem
	for _, modelRecord := range modelRecords {
		model := dtos.NewModelItemFromRecord(modelRecord)
		models = append(models, model)

		if !seenMakeIDs[model.MakeID] {
			seenMakeIDs[model.MakeID] = true
			makeIDs = append(makeIDs, model.MakeID)
		}
	}

	// Populating makes' data.
	makeRecords, err := makedata.GetMakesByIDs(context.Session, makeIDs)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	makesByIDs := make(map[string]*makedata.MakeRecord)
	for _, m := range makeRecords {
		makesByIDs[m.GetIDFromTypeAndID()] = m
	}

	for _, m := range models {
		makeRecord := makesByIDs[m.MakeID]
		if makeRecord != nil {
			m.Make = makeRecord.Name
			m.MakeImage = makeRecord.Image
		}
	}

	response := getModelsPage(context, models)
	dtos.GetCommentsForModels(context.Session, response.Models)

	return httpresponses.CreateJSONResponse(200, response), nil
}

func getModelsPage(context requestcontext.RequestContext, models []*dtos.ModelItem) dtos.ModelList {
	// Pagination
	page := context.APIRequest.QueryStringParameters["page"]
	pageIndex, err := strconv.Atoi(page)
	if err != nil {
		pageIndex = 1
	}

	orderBy := context.APIRequest.QueryStringParameters["orderBy"]

	modelList := dtos.GetModelsPage(models, pageIndex, defaultPageSize, orderBy)

	return modelList
}
