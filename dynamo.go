package main

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

const tableName = "debris_data"

func GetPlasticFromDate(date string, 
	dynaClient dynamodbiface.DynamoDBAPI) (*PlasticJson, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"date": {
				S: aws.String(date),
			},
		},
	}

	result, err := dynaClient.GetItem(input)
	if err != nil {
		return nil, errors.New("Error Failed to fetch record!")
	}

	item := new(PlasticJson)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, errors.New("error failed to unmarshall dynamo record")
	}

	return item, nil
}

func createDynamoSession() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(
		session.Options{
			SharedConfigState: session.SharedConfigEnable,
		},
	))

	return dynamodb.New(sess)
}
