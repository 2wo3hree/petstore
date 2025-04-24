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
	IssueBook(ctx context.Context, userID, bookID int) (models.Rental, error)
	ReturnBook(ctx context.Context, userID, bookID int) error
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

func (s *bookService) IssueBook(ctx context.Context, userID, bookID int) (models.Rental, error) {
	return s.repo.IssueBook(ctx, userID, bookID)
}

func (s *bookService) ReturnBook(ctx context.Context, userID, bookID int) error {
	return s.repo.ReturnBook(ctx, userID, bookID)
}
