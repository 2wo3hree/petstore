package repository

import (
	"context"
	"petstore/internal/models"
)

type OrderRepository interface {
	Create(ctx context.Context, order models.Order) error
	GetByID(ctx context.Context, id int64) (models.Order, error)
	Delete(ctx context.Context, id int64) error
	GetInventory(ctx context.Context) (map[string]int, error)
}
