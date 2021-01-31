package dashboard

import (
	"log"
	"strings"

	"buggy/api/requestcontext"
	"buggy/internal/data/makedata"
	"buggy/internal/data/modeldata"
	"buggy/internal/httpresponses"

	"github.com/aws/aws-lambda-go/events"
)

type dashboardMake struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Votes int    `json:"votes"`
}

type dashboardModel struct {
	ID    string `json:"id"`
	Make  string `json:"make"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Votes int    `json:"votes"`
}

// DashboardResponse is a DTO returned by the dashboard handler.
type DashboardResponse struct {
	Make  dashboardMake  `json:"make"`
	Model dashboardModel `json:"model"`
}

// Handler handles the dashboard-related API requests.
func Handler(context requestcontext.RequestContext) (events.APIGatewayProxyResponse, error) {
	switch {
	case strings.EqualFold(context.APIRequest.HTTPMethod, "GET") && context.Path == "dashboard":
		return dashboardHandler(context)
	}

	return httpresponses.NotFound, nil
}

func dashboardHandler(context requestcontext.RequestContext) (events.APIGatewayProxyResponse, error) {
	topMake, err := makedata.GetTopMake(context.Session)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	topModel, err := modeldata.GetTopModel(context.Session)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	var response DashboardResponse
	if topMake != nil {
		response.Make = dashboardMake{
			ID:    topMake.GetIDFromTypeAndID(),
			Name:  topMake.Name,
			Image: topMake.Image,
			Votes: topMake.Votes,
		}
	}

	if topModel != nil {
		// Fetch the make
		make, err := makedata.GetMakeByID(context.Session, topModel.GetMakeID())
		if err != nil {
			log.Printf("Unable to fetch model's make: %v\n", err)
		}

		var makeName string
		if make != nil {
			makeName = make.Name
		}

		response.Model = dashboardModel{
			ID:    topModel.GetIDFromTypeAndID(),
			Make:  makeName,
			Name:  topModel.Name,
			Image: topModel.Image,
			Votes: topModel.Votes,
		}
	}

	return httpresponses.CreateJSONResponse(200, response), nil
}
