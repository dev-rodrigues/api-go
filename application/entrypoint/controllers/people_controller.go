package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api/application/config/database"
	"rest-api/application/repositories/impl"
	"rest-api/application/services"
)

var (
	db, _   = database.DbConnect()
	service = services.NewPeopleService(impl.NewPeopleRepository(db))
)

func GetPeoples(w http.ResponseWriter, r *http.Request) {
	response, err := service.GetPeoples()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&response)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {

}

func CreatePerson(w http.ResponseWriter, r *http.Request) {

}

func DeletePerson(w http.ResponseWriter, r *http.Request) {

}
