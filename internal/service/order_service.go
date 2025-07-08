// internal/service/order_service.go
package service

import (
    "context"
    "github.com/google/uuid"
    "github.com/paragon-corp/backend-paragon/internal/model"
    "github.com/paragon-corp/backend-paragon/internal/storage"
)

func CreateOrder(coffeeID string) (model.Order, error) {
	order := model.Order{
		ID:       uuid.New().String(),
		CoffeeID: coffeeID,
		Status:   "pending",
	}

	_, err := storage.WriteDB.Exec(context.Background(),
		"INSERT INTO orders (id, coffee_id, status) VALUES ($1, $2, $3)",
		order.ID, order.CoffeeID, order.Status)

	return order, err
}
func GetOrder(id string) (model.Order, error) {
	var order model.Order

	row := storage.WriteDB.QueryRow(context.Background(),
		"SELECT id, coffee_id, status FROM orders WHERE id = $1", id)

	err := row.Scan(&order.ID, &order.CoffeeID, &order.Status)
	return order, err
}