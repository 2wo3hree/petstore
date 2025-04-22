package service

import (
	"context"
	"errors"
	"fmt"
	"petstore/internal/models"
	"petstore/internal/repository"
)

type LibrarySuperService struct {
	users    UserService
	books    BookService
	authors  AuthorService
	rentRepo repository.RentalRepository
}

func NewLibrarySuperService(u UserService, b BookService, a AuthorService, rent repository.RentalRepository) *LibrarySuperService {
	return &LibrarySuperService{
		users:    u,
		books:    b,
		authors:  a,
		rentRepo: rent,
	}
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
	rental, err := s.rentRepo.IssueBook(ctx, userID, bookID)
	if err != nil {
		if errors.Is(err, repository.ErrAlreadyIssued) {
			return models.Rental{}, fmt.Errorf("book already issued")
		}
		return models.Rental{}, fmt.Errorf("failed to issue: %w", err)
	}

	return rental, nil
}

func (s *LibrarySuperService) Return(ctx context.Context, userID, bookID int) error {
	err := s.rentRepo.ReturnBook(ctx, userID, bookID)
	if err != nil {
		return fmt.Errorf("return error: %w", err)
	}

	return nil
}

func (s *LibrarySuperService) GetTopAuthors(ctx context.Context) ([]models.AuthorCount, error) {
	return s.authors.TopAuthors(ctx)
}
