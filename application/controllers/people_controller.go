package controllers

import (
	"log"
	"net/http"
	"rest-api/application/controllers/dto"
)

type ManyPeoples []dto.Person

func GetPeople(w http.ResponseWriter, r *http.Request) {
	log.Println("cuco")
}

func GetPerson(w http.ResponseWriter, r *http.Request) {

}

func CreatePerson(w http.ResponseWriter, r *http.Request) {

}

func DeletePerson(w http.ResponseWriter, r *http.Request) {

}
