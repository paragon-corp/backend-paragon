package handler

import (
	"encoding/json"
	"net/http"
	"github.com/paragon-corp/backend-paragon/internal/model"
	"github.com/paragon-corp/backend-paragon/internal/service"
)

func GetCoffees(w http.ResponseWriter, r *http.Request) {
	coffees, err := service.ListCoffees()
	if err != nil {
		http.Error(w, "Failed to fetch coffees", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(coffees)
}
func CreateCoffee(w http.ResponseWriter, r *http.Request) {
	var coffee model.Coffee

	if err := json.NewDecoder(r.Body).Decode(&coffee); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := service.CreateCoffee(coffee)
	if err != nil {
		http.Error(w, "Failed to create coffee", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}