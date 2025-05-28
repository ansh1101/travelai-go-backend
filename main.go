package main

import (
	"fmt"
	"log"
	"net/http"
	"travel-api/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/accommodations", handlers.GetAccommodations).Methods("GET")
	r.HandleFunc("/accommodations", handlers.CreateAccommodation).Methods("POST")
	r.HandleFunc("/accommodations/{id}", handlers.GetAccommodationByID).Methods("GET")
	r.HandleFunc("/accommodations/{id}", handlers.DeleteAccommodation).Methods("DELETE")
	r.HandleFunc("/accommodations/{id}", handlers.UpdateAccommodation).Methods("PUT")

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
