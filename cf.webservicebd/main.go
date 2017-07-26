package main

import (
	"log"
	"net/http"

	"./handlers"
	"github.com/julienschmidt/httprouter"
)

func main() {

	mux := httprouter.New()
	mux.GET("/api/products", handlers.GetAllProducts)
	mux.GET("/api/products/:id", handlers.GetProduct)
	// mux.GET("/api/users", handlers.CreateProduct).Methods("POST")
	// mux.HandleFunc("/api/users/{id:[0-9]+}", handlers.UpdateProductr).Methods("PUT")
	// mux.HandleFunc("/api/users/{id:[0-9]+}", handlers.DeleteProduct).Methods("DELETE")

	log.Println("El servidor esta a la escucha 80")
	log.Fatal(http.ListenAndServe(":80", mux))
}
