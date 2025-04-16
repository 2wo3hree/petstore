package repository

import (
	"context"
	"petstore/internal/models"
)

type PetRepository interface {
	Create(ctx context.Context, pet models.Pet) error
	GetByID(ctx context.Context, id int64) (models.Pet, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, pet models.Pet) error
	List(ctx context.Context, status string) ([]models.Pet, error)
}
