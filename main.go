package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	dynaClient dynamodbiface.DynamoDBAPI
)

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func response(body interface{}, statusCode int) events.APIGatewayProxyResponse {
	resp := events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		},
	}

	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)
	return resp
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	rawDate, found := req.PathParameters["date"]
	if !found {
		fmt.Println("error, no date parameter found in path param")
		return response("no date path paramter supplied", http.StatusBadRequest), nil
	}
	date, err := url.QueryUnescape(rawDate)
	if err != nil {
		return response(ErrorBody{
			aws.String(err.Error()),
		}, http.StatusInternalServerError), err
	}

	plasticItem, err := GetPlasticFromDate(date, dynaClient)
	if err != nil {
		fmt.Errorf("error trying to get plastic from dynamo")
		return response(ErrorBody{
              aws.String(err.Error()),
          }, http.StatusInternalServerError), err
	}

	return response(plasticItem, http.StatusOK), nil
}

func main() {
	dynaClient = createDynamoSession()
	fmt.Println("cold starting lambda..")
	lambda.Start(Handler)
}
