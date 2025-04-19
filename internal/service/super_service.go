package service

import (
	"context"
	"fmt"
	"petstore/internal/models"
)

type LibrarySuperService struct {
	users   UserService
	books   BookService
	rentals RentalService
	authors AuthorService
}

func NewLibrarySuperService(u UserService, b BookService, r RentalService, a AuthorService) *LibrarySuperService {
	return &LibrarySuperService{
		users:   u,
		books:   b,
		rentals: r,
		authors: a,
	}
}

func (s *LibrarySuperService) Issue(ctx context.Context, userID, bookID int) (models.Rental, error) {
	user, err := s.users.GetByID(ctx, userID)
	if err != nil {
		return models.Rental{}, fmt.Errorf("user not found: %w", err)
	}

	book, err := s.books.GetByID(ctx, bookID)
	if err != nil {
		return models.Rental{}, fmt.Errorf("book not found: %w", err)
	}

	rental, err := s.rentals.IssueBook(ctx, user.ID, book.ID)
	if err != nil {
		return models.Rental{}, fmt.Errorf("failed to issue: %w", err)
	}

	return rental, nil
}

func (s *LibrarySuperService) Return(ctx context.Context, userID, bookID int) error {
	err := s.rentals.ReturnBook(ctx, userID, bookID)
	if err != nil {
		return fmt.Errorf("return error: %w", err)
	}

	return nil
}

func (s *LibrarySuperService) GetTopAuthors(ctx context.Context) ([]models.AuthorCount, error) {
	return s.authors.TopAuthors(ctx)
}
