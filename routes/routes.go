package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"rest-api/application/controllers"
	"time"
)

func Route() {
	router := mux.NewRouter()

	router.HandleFunc("/contact", controllers.GetPeople).Methods("GET")
	router.HandleFunc("/contact/{id}", controllers.GetPerson).Methods("GET")
	router.HandleFunc("/contact/{id}", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/contact/{id}", controllers.DeletePerson).Methods("DELETE")

	host := fmt.Sprintf(viper.GetString("app.host"))

	srv := &http.Server{
		Handler:      router,
		Addr:         host,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
