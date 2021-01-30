package httpresponses

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

var NotFound = CreateAPIResponse(404, []byte("Not Found"))
var InvalidRequest = CreateAPIResponse(400, []byte("Unable to decode the request."))

// CreateJSONResponse creates a standard JSON data response.
func CreateJSONResponse(statusCode int, data interface{}) events.APIGatewayProxyResponse {
	var bodyJSON []byte
	var err error
	bodyJSON, err = json.Marshal(data)
	if err != nil {
		log.Fatalf("Unable to serialize response value: %v. Error: %v\n", data, err)
		return CreateAPIResponse(500, []byte("An unexpected error has occurred"))
	}

	return CreateAPIResponse(200, bodyJSON)
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
