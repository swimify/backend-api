package main

import (
	"log"
	"net/http"
	"github.com/swimify/backend-api/route"

	"github.com/gorilla/mux"
)



func main() {
	log.Println("starting API server")
	//create a new router
	router := mux.NewRouter()
	log.Println("creating routes")
	//specify endpoints
	router.HandleFunc("/health-check", route.HealthCheck).Methods("GET")
	router.HandleFunc("/persons", route.Persons).Methods("GET")
	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)

}



