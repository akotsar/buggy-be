package httpresponses

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

var NotFound = CreateErrorResponse(404, "Not Found")
var InvalidRequest = CreateErrorResponse(400, "Invalid request.")
var Unauthorized = CreateErrorResponse(401, "Not authorized.")

type errorResponse struct {
	Message string `json:"message"`
}

// CreateJSONResponse creates a standard JSON data response.
func CreateJSONResponse(statusCode int, data interface{}) events.APIGatewayProxyResponse {
	var bodyJSON []byte
	var err error
	bodyJSON, err = json.Marshal(data)
	if err != nil {
		log.Printf("Unable to serialize response value: %v. Error: %v\n", data, err)
		return CreateAPIResponse(500, []byte("An unexpected error has occurred"))
	}

	return CreateAPIResponse(statusCode, bodyJSON)
}

// CreateAPIResponse creates a standard API response.
func CreateAPIResponse(statusCode int, body []byte) events.APIGatewayProxyResponse {
	headers := map[string]string{
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Allow-Credentials": "true",
		"Content-Type":                     "application/json",
	}

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(body),
		Headers:    headers}
}

// CreateErrorResponse creates a standard error response.
func CreateErrorResponse(statusCode int, errorMessage string) events.APIGatewayProxyResponse {
	return CreateJSONResponse(statusCode, &errorResponse{Message: errorMessage})
}
