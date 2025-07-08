// internal/handler/order_handler.go
package handler

import (
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/paragon-corp/backend-paragon/internal/service"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var input struct {
		CoffeeID string `json:"coffee_id"`
	}
	_ = json.NewDecoder(r.Body).Decode(&input)

	order, err := service.CreateOrder(input.CoffeeID) // ✅ correct unpack
	if err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(order)
}


func GetOrder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	order, err := service.GetOrder(id)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(order)
}

