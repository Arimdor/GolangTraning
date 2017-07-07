package main

import (
	"log"
	"net/http"

	"./handlers"
	"./models"

	"github.com/gorilla/mux"
)

func main() {

	mux := mux.NewRouter()
	models.SetDefaultUser()
	mux.HandleFunc("/api/users", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/users", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/users/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/users/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	log.Println("El servidor esta a la escucha 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
//gg