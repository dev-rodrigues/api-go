package controllers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"rest-api/application/config/database"
	database2 "rest-api/application/datasources/database"
	"rest-api/application/entrypoint/controllers/dto"
	"rest-api/application/entrypoint/controllers/mappers"
	"rest-api/application/services"
	"rest-api/application/utils"
)

var (
	db, _   = database.DbConnect()
	service = services.NewPeopleService(database2.NewPeopleRepository(db))
)

func GetPeoples(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response, err := service.GetPeoples()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := utils.Map(response, mappers.PeopleToPeopleDto)

	json.NewEncoder(w).Encode(&result)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, error := service.GetPeople(id)

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := mappers.PeopleToPeopleDto(response)

	json.NewEncoder(w).Encode(&result)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var peopleDto dto.PersonDTO
	peopleDto.ID = uuid.New().String()

	err := json.NewDecoder(r.Body).Decode(&peopleDto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	people := mappers.PeopleDTOToPeople(&peopleDto)

	response, error := service.CreatePeople(people)

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := mappers.PeopleToPeopleDto(response)

	json.NewEncoder(w).Encode(result)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := service.DeletePeople(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
