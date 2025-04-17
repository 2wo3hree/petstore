package service

import (
	"context"
	"petstore/internal/models"
)

type LibraryService interface {
	IssueBook(ctx context.Context, userID, bookID int) (models.Rental, error)
	ReturnBook(ctx context.Context, userID, bookID int) error
	GetTopAuthors(ctx context.Context) ([]models.AuthorCount, error)
}

type libraryService struct {
	users   UserService
	books   BookService
	rentals RentalService
	authors AuthorService
}

func NewLibraryService(u UserService, b BookService, r RentalService, a AuthorService) LibraryService {
	return &libraryService{
		users:   u,
		books:   b,
		rentals: r,
		authors: a,
	}
}

func (l *libraryService) IssueBook(ctx context.Context, userID, bookID int) (models.Rental, error) {
	return l.rentals.IssueBook(ctx, userID, bookID)
}

func (s *libraryService) ReturnBook(ctx context.Context, userID, bookID int) error {
	return s.rentals.ReturnBook(ctx, userID, bookID)
}

func (s *libraryService) GetTopAuthors(ctx context.Context) ([]models.AuthorCount, error) {
	return s.authors.TopAuthors(ctx)
}
