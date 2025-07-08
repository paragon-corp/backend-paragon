// cmd/api/main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/paragon-corp/backend-paragon/internal/handler"
	"github.com/paragon-corp/backend-paragon/internal/storage"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found or failed to load it")
    }

	storage.InitDB()
	fmt.Println("DB Initialized")
	defer storage.ReadDB.Close()
	defer storage.WriteDB.Close()
	fmt.Println("DB Deferred")
	r := chi.NewRouter()
	fmt.Println("Router Created")

	storage.InitSchema()
	fmt.Println("Schema Initialized")


	r.Get("/coffees", handler.GetCoffees)
	r.Post("/coffees", handler.CreateCoffee)
	r.Post("/orders", handler.CreateOrder)
	r.Get("/orders/{id}", handler.GetOrder)

	log.Println("Coffee service running on :8080")
	http.ListenAndServe(":8080", r)
}
