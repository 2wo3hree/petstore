package repository

import (
	"context"
	"petstore/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user models.User) (int64, error)
	GetByID(ctx context.Context, id int) (models.User, error)
	List(ctx context.Context, limit, offset int) ([]models.User, int, error)
	GetWithRentals(ctx context.Context, id int) (models.User, error)
}
