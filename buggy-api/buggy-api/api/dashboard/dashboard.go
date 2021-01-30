package dashboard

import (
	"log"
	"strings"

	"buggy/internal/data/makedata"
	"buggy/internal/data/modeldata"
	"buggy/internal/httpresponses"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/session"
)

type dashboardMake struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Votes int    `json:"votes"`
}

type dashboardModel struct {
	Id    string `json:"id"`
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

// Handler handles the user-related API requests.
func Handler(userID string, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch {
	case strings.EqualFold(request.HTTPMethod, "GET") && request.PathParameters["thepath"] == "dashboard":
		return dashboardHandler(userID)
	}

	return httpresponses.NotFound, nil
}

func dashboardHandler(userID string) (events.APIGatewayProxyResponse, error) {
	session := session.Must(session.NewSession())

	topMake, err := makedata.GetTopMake(session)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	topModel, err := modeldata.GetTopModel(session)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	var response DashboardResponse
	if topMake != nil {
		response.Make = dashboardMake{
			Id:    topMake.GetIDFromTypeAndID(),
			Name:  topMake.Name,
			Image: topMake.Image,
			Votes: topMake.Votes,
		}
	}

	if topModel != nil {
		// Fetch the make
		make, err := makedata.GetMakeByID(session, topModel.ModelID)
		if err != nil {
			log.Fatalf("Unable to fetch model's make: %v\n", err)
		}

		var makeName string
		if make != nil {
			makeName = make.Name
		}

		response.Model = dashboardModel{
			Id:    topModel.GetIDFromTypeAndID(),
			Make:  makeName,
			Name:  topModel.Name,
			Image: topModel.Image,
			Votes: topModel.Votes,
		}
	}

	return httpresponses.CreateJSONResponse(200, response), nil
}
