package service

import (
	"context"

	"github.com/paragon-corp/backend-paragon/internal/model"
	"github.com/paragon-corp/backend-paragon/internal/storage"
)

func ListCoffees() ([]model.Coffee, error) {
	rows, err := storage.ReadDB.Query(context.Background(), "SELECT id, name, price FROM coffees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var coffees []model.Coffee
	for rows.Next() {
		var c model.Coffee
		if err := rows.Scan(&c.ID, &c.Name, &c.Price); err != nil {
			return nil, err
		}
		coffees = append(coffees, c)
	}
	return coffees, nil
}

func CreateCoffee(c model.Coffee) error {
	_, err := storage.WriteDB.Exec(context.Background(),
		"INSERT INTO coffees (id, name, price) VALUES ($1, $2, $3)",
		c.ID, c.Name, c.Price)
	return err
}
