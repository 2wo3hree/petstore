package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"petstore/internal/models"
	"petstore/internal/repository"
)

type orderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) repository.OrderRepository {
	return &orderRepo{db: db}
}

func (r *orderRepo) Create(ctx context.Context, order models.Order) error {
	query := `INSERT INTO orders (id, pet_id, quantity, ship_date, status, complete)
			  VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(ctx, query, order.ID, order.PetID, order.Quantity, order.ShipDate, order.Status, order.Complete)
	return err
}

func (r *orderRepo) GetByID(ctx context.Context, id int64) (models.Order, error) {
	query := `SELECT id, pet_id, quantity, ship_date, status, complete FROM orders WHERE id=$1`
	row := r.db.QueryRow(ctx, query, id)

	var order models.Order
	err := row.Scan(&order.ID, &order.PetID, &order.Quantity, &order.ShipDate, &order.Status, &order.Complete)
	return order, err
}

func (r *orderRepo) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM orders WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *orderRepo) GetInventory(ctx context.Context) (map[string]int, error) {
	query := `
		SELECT status, COUNT(*) 
		FROM orders 
		GROUP BY status;
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	inventory := make(map[string]int)
	for rows.Next() {
		var status string
		var count int
		if err := rows.Scan(&status, &count); err != nil {
			return nil, err
		}
		inventory[status] = count
	}

	return inventory, nil
}
