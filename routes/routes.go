package routes

import (
	"github.com/gorilla/mux"
)

func AttachRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/persons", Persons).Methods("GET")
	return router
}
