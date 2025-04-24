package service

import (
	"context"
	"errors"
	"fmt"
	"petstore/internal/models"
	"petstore/internal/repository"
)

type Facade interface {
	Issue(ctx context.Context, userID, bookID int) (models.Rental, error)
	Return(ctx context.Context, userID, bookID int) error
	GetTopAuthors(ctx context.Context) ([]models.AuthorCount, error)
	CreateUser(ctx context.Context, u models.User) (int64, error)
	ListUsers(ctx context.Context, limit int, offset int) ([]models.User, int, error)
	GetUserByID(ctx context.Context, id int) (models.User, error)
	GetWithRentals(ctx context.Context, id int) (models.User, error)
	CreateBook(ctx context.Context, b models.Book) (int64, error)
	GetBookByID(ctx context.Context, id int) (models.Book, error)
	ListBooks(ctx context.Context, limit int, offset int) ([]models.Book, int, error)
	CreateAuthor(ctx context.Context, a models.Author) (int64, error)
	GetAuthorByID(ctx context.Context, id int64) (models.Author, error)
	ListAuthors(ctx context.Context, limit int, offset int) ([]models.Author, int, error)
}

type LibrarySuperService struct {
	users   UserService
	books   BookService
	authors AuthorService
}

func NewLibrarySuperService(u UserService, b BookService, a AuthorService) *LibrarySuperService {
	return &LibrarySuperService{
		users:   u,
		books:   b,
		authors: a,
	}
}

func (s *LibrarySuperService) CreateUser(ctx context.Context, u models.User) (int64, error) {
	return s.users.Create(ctx, u)
}

func (s *LibrarySuperService) ListUsers(ctx context.Context, limit int, offset int) ([]models.User, int, error) {
	return s.users.List(ctx, limit, offset)
}

func (s *LibrarySuperService) GetUserByID(ctx context.Context, id int) (models.User, error) {
	return s.users.GetByID(ctx, id)
}

func (s *LibrarySuperService) GetWithRentals(ctx context.Context, id int) (models.User, error) {
	return s.users.GetWithRentals(ctx, id)
}

func (s *LibrarySuperService) CreateBook(ctx context.Context, b models.Book) (int64, error) {
	return s.books.Create(ctx, b)
}

func (s *LibrarySuperService) GetBookByID(ctx context.Context, id int) (models.Book, error) {
	return s.books.GetByID(ctx, id)
}

func (s *LibrarySuperService) ListBooks(ctx context.Context, limit int, offset int) ([]models.Book, int, error) {
	return s.books.List(ctx, limit, offset)
}

func (s *LibrarySuperService) CreateAuthor(ctx context.Context, a models.Author) (int64, error) {
	return s.authors.Create(ctx, a)
}

func (s *LibrarySuperService) GetAuthorByID(ctx context.Context, id int64) (models.Author, error) {
	return s.authors.GetByID(ctx, id)
}

func (s *LibrarySuperService) ListAuthors(ctx context.Context, limit int, offset int) ([]models.Author, int, error) {
	return s.authors.List(ctx, limit, offset)
}

func (s *LibrarySuperService) Issue(ctx context.Context, userID, bookID int) (models.Rental, error) {
	_, err := s.users.GetByID(ctx, userID)
	if err != nil {
		return models.Rental{}, fmt.Errorf("user not found: %w", err)
	}

	_, err = s.books.GetByID(ctx, bookID)
	if err != nil {
		return models.Rental{}, fmt.Errorf("book not found: %w", err)
	}

	// проверка, свободна ли книга, и сама аренда
	rental, err := s.books.IssueBook(ctx, userID, bookID)
	if err != nil {
		if errors.Is(err, repository.ErrAlreadyIssued) {
			return models.Rental{}, fmt.Errorf("book already issued")
		}
		return models.Rental{}, fmt.Errorf("failed to issue: %w", err)
	}

	return rental, nil
}

func (s *LibrarySuperService) Return(ctx context.Context, userID, bookID int) error {
	err := s.books.ReturnBook(ctx, userID, bookID)
	if err != nil {
		return fmt.Errorf("return error: %w", err)
	}

	return nil
}

func (s *LibrarySuperService) GetTopAuthors(ctx context.Context) ([]models.AuthorCount, error) {
	return s.authors.TopAuthors(ctx)
}
