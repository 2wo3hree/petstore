package repository

import (
	"context"
	"petstore/internal/models"
)

type AuthorRepository interface {
	Create(ctx context.Context, a models.Author) (int64, error)
	GetByID(ctx context.Context, id int64) (models.Author, error)
	List(ctx context.Context, limit, offset int) ([]models.Author, int, error)
	TopAuthors(ctx context.Context) ([]models.AuthorCount, error)
}
