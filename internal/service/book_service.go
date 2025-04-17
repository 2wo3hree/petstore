package service

import (
	"context"
	"petstore/internal/models"
	"petstore/internal/repository"
)

type BookService interface {
	Create(ctx context.Context, b models.Book) (int64, error)
	GetByID(ctx context.Context, id int) (models.Book, error)
	List(ctx context.Context, limit, offset int) ([]models.Book, int, error)
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(r repository.BookRepository) BookService {
	return &bookService{repo: r}
}

func (s *bookService) Create(ctx context.Context, b models.Book) (int64, error) {
	return s.repo.Create(ctx, b)
}

func (s *bookService) GetByID(ctx context.Context, id int) (models.Book, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *bookService) List(ctx context.Context, limit, offset int) ([]models.Book, int, error) {
	return s.repo.List(ctx, limit, offset)
}
