package main

import (
	"encoding/json"
)

type Message struct {
    Name string
    Body string
    Time int64
}

func retrieveDynamoJsonInfo(date string) []byte {
	//TODO: integrate with dynamodb to retrieve debris data for this date
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m)
	if err != nil {
		panic("AHH! error marshalling!")
	}

	return b
}