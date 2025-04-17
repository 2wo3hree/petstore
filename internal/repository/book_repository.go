package repository

import (
	"context"
	"petstore/internal/models"
)

type BookRepository interface {
	Create(ctx context.Context, b models.Book) (int64, error)
	GetByID(ctx context.Context, id int) (models.Book, error)
	List(ctx context.Context, limit, offset int) ([]models.Book, int, error)
}
