// internal/model/models.go
package model

type Coffee struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Order struct {
	ID       string  `json:"id"`
	CoffeeID string  `json:"coffee_id"`
	Status   string  `json:"status"` // e.g., "pending", "brewing", "done"
}
