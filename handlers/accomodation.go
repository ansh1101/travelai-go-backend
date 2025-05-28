package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"travel-api/models"

	"github.com/gorilla/mux"
)

var accommodations []models.Accommodation
var idCounter = 1

func GetAccommodations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accommodations)
}

func CreateAccommodation(w http.ResponseWriter, r *http.Request) {
	var acc models.Accommodation
	err := json.NewDecoder(r.Body).Decode(&acc)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	acc.ID = idCounter
	idCounter++

	accommodations = append(accommodations, acc)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(acc)
}

func GetAccommodationByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, acc := range accommodations {
		if acc.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(acc)
			return
		}
	}
	http.Error(w, "Accommodation not found", http.StatusNotFound)

}

func UpdateAccommodation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid msg", http.StatusBadRequest)
	}

	var updateAcc models.Accommodation
	err = json.NewDecoder(r.Body).Decode(&updateAcc)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i, acc := range accommodations {

		if acc.ID == id {
			accommodations[i].Name = updateAcc.Name
			accommodations[i].Location = updateAcc.Location
			accommodations[i].Price = updateAcc.Price

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(accommodations[i])
			return
		}
	}
	http.Error(w, "Accommodation not found", http.StatusNotFound)

}
func DeleteAccommodation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for index, acc := range accommodations {
		if acc.ID == id {
			fmt.Println(acc, id)
			// Remove the item from the slice
			accommodations = append(accommodations[:index], accommodations[index+1:]...)

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Accommodation deleted"))
			return
		}
	}

	http.Error(w, "Accommodation not found", http.StatusNotFound)
}
