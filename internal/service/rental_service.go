package service

import (
	"context"
	"petstore/internal/models"
	"petstore/internal/repository"
)

type RentalService interface {
	IssueBook(ctx context.Context, userID, bookID int) (models.Rental, error)
	ReturnBook(ctx context.Context, userID, bookID int) error
}

type rentalService struct {
	rRepo repository.RentalRepository
	uRepo repository.UserRepository
	bRepo repository.BookRepository
}

func NewRentalService(r repository.RentalRepository, u repository.UserRepository, b repository.BookRepository) RentalService {
	return &rentalService{rRepo: r, uRepo: u, bRepo: b}
}

func (s *rentalService) IssueBook(ctx context.Context, userID, bookID int) (models.Rental, error) {
	return s.rRepo.IssueBook(ctx, userID, bookID)
}

func (s *rentalService) ReturnBook(ctx context.Context, userID, bookID int) error {
	return s.rRepo.ReturnBook(ctx, userID, bookID)
}
