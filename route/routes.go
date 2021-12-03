package route


import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/swimify/backend-api/model"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running  woot")
}

func Persons(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	var persons []model.Person

	var person model.Person
	person.Id = 1
	person.FirstName = "Michael"
	person.LastName = "Phelps"
	persons = append(persons, person)

	person.Id = 2
	person.FirstName = "Dana"
	person.LastName = "Torres"
	persons = append(persons, person)

	person.Id = 3
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