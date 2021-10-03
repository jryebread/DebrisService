package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func response(body string, statusCode int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse {
		StatusCode: statusCode,
		Body: string(body),
		Headers: map[string]string {
			"Access-Control-Allow-Origin": "*",
		},
	}
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	rawDate, found := req.PathParameters["date"]
	if !found {
		fmt.Println("error, no date parameter found in path param")
		return response("no date path paramter supplied", http.StatusBadRequest), nil
	}
	date, err := url.QueryUnescape(rawDate)
	if err != nil {
		return response("error trying to query unescape", http.StatusInternalServerError), err
	}
	
	jsonData :=	retrieveDynamoJsonInfo(date)

	return response(string(jsonData), http.StatusOK), nil
}

func main() {
	lambda.Start(Handler)
}
