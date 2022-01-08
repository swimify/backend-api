package main

import (
	"log"
	"net/http"

	"github.com/swimlibre/backend-api/container"
	"github.com/swimlibre/backend-api/fixtures"
)

func main() {
	// Instantiate the DI Container
	container := container.BuildContainer()

	// @todo move this to a CLI command
	// Load Fixtures
	fixtures.LoadFixtures(container.DB)

	//start and listen to requests
	log.Println("listening on 8080")
	http.ListenAndServe(":8080", container.Router)
}
