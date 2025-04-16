package repository

import (
	"context"
	"petstore/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user models.User) error
	GetByUsername(ctx context.Context, username string) (models.User, error)
	Update(ctx context.Context, username string, updated models.User) error
	Delete(ctx context.Context, username string) error
}
