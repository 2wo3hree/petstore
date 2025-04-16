package service

import (
	"context"
	"petstore/internal/models"
	"petstore/internal/repository"
)

type OrderService interface {
	Create(ctx context.Context, order models.Order) error
	GetByID(ctx context.Context, id int64) (models.Order, error)
	Delete(ctx context.Context, id int64) error
	GetInventory(ctx context.Context) (map[string]int, error)
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(r repository.OrderRepository) OrderService {
	return &orderService{repo: r}
}

func (s *orderService) Create(ctx context.Context, order models.Order) error {
	return s.repo.Create(ctx, order)
}

func (s *orderService) GetByID(ctx context.Context, id int64) (models.Order, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *orderService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
func (s *orderService) GetInventory(ctx context.Context) (map[string]int, error) {
	return s.repo.GetInventory(ctx)
}
