package main

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

const tableName = "plastic-debris-table"

func GetAllDatesFromTable(dynaClient dynamodbiface.DynamoDBAPI) (*DateResponse, error) {
	proj := expression.NamesList(expression.Name("date"))
	expr, err := expression.NewBuilder().WithProjection(proj).Build()
	if err != nil {
		fmt.Printf("error trying to build proj %s", err.Error())
		return nil, err
	}
	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		TableName:                 aws.String(tableName),
		ProjectionExpression:      expr.Projection(),
	}

	result, err := dynaClient.Scan(input)
	if err != nil {
		fmt.Printf("error trying to scan dynamo with input %s", err.Error())
		return nil, err
	}

	num_items := 0
	retVal := DateResponse{}	
	for _, i := range result.Items {
		item := PlasticJson{}
		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Printf("error trying to unmarshal dynamo entry %s", err.Error())
			return nil, err
		}

		num_items += 1
		retVal.Dates = append(retVal.Dates, item.Date)
	}
	fmt.Println("Found this many items for scan:", num_items)
	return &retVal, nil
}

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
