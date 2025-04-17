package service

import (
	"context"
	"petstore/internal/models"
	"petstore/internal/repository"
)

type AuthorService interface {
	Create(ctx context.Context, a models.Author) (int64, error)
	GetByID(ctx context.Context, id int64) (models.Author, error)
	List(ctx context.Context, limit, offset int) ([]models.Author, int, error)
	TopAuthors(ctx context.Context) ([]models.AuthorCount, error)
}

type authorService struct {
	repo repository.AuthorRepository
}

func NewAuthorService(r repository.AuthorRepository) AuthorService {
	return &authorService{repo: r}
}

func (s *authorService) Create(ctx context.Context, a models.Author) (int64, error) {
	return s.repo.Create(ctx, a)
}
func (s *authorService) GetByID(ctx context.Context, id int64) (models.Author, error) {
	return s.repo.GetByID(ctx, id)
}
func (s *authorService) List(ctx context.Context, limit, offset int) ([]models.Author, int, error) {
	return s.repo.List(ctx, limit, offset)
}
func (s *authorService) TopAuthors(ctx context.Context) ([]models.AuthorCount, error) {
	return s.repo.TopAuthors(ctx)
}
