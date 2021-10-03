package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func debrisHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	date, ok := param["date"]
	if !ok {
		panic("Error trying to parse date from vars in request")
	}
	jsonData := retrieveDynamoJsonInfo(date)

	fmt.Println(string(jsonData))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
	
}

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/", homePage)
	route.HandleFunc("/plasticinfo/{date}", debrisHandler)
	http.ListenAndServe(":8000", route)
}
