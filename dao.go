package main

import (
)

type PlasticClusterData struct {
	Lat string `json:"lat"`
	Long string `json:"long"`
	Timestamp int `json:"timestamp"`
	Description string `json:"description"` // or size?
}

type PlasticJson struct {
	Date string `json:"date"`
	PlasticClusterDataArr []PlasticClusterData `json:"plastic_cluster_data"`
}

type DateResponse struct {
	Dates []string `json:"dates`
}
type Message struct {
    Name string
    Body string
    Time int64
}

// func retrieveDynamoJsonInfo(date string) []byte {
// 	//TODO: integrate with dynamodb to retrieve debris data for this date
// 	m := Message{"Alice", "Hello", 1294706395881547000}
// 	b, err := json.Marshal(m)
// 	if err != nil {
// 		panic("AHH! error marshalling!")
// 	}

// 	return b
// }