package routes

import (
	"encoding/json"
	"net/http"

	"github.com/swimlibre/backend-api/model"
)

func Persons(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	var persons []model.Person

	var person model.Person
	person.ID = 1
	person.FirstName = "Michael"
	person.LastName = "Phelps"
	persons = append(persons, person)

	person.ID = 2
	person.FirstName = "Dana"
	person.LastName = "Torres"
	persons = append(persons, person)

	person.ID = 3
	person.FirstName = "Mark"
	person.LastName = "Spitz"
	persons = append(persons, person)

	response.Persons = persons

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}
