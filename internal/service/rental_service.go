package service

import (
	"context"
	"errors"
	"petstore/internal/models"
	"petstore/internal/repository"
)

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrBookNotFound  = errors.New("book not found")
	ErrAlreadyIssued = repository.ErrAlreadyIssued
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
	// Проверяем, что пользователь есть
	if _, err := s.uRepo.GetByID(ctx, userID); err != nil {
		return models.Rental{}, ErrUserNotFound
	}
	// Проверяем, что книга есть
	if _, err := s.bRepo.GetByID(ctx, bookID); err != nil {
		return models.Rental{}, ErrBookNotFound
	}
	// Пытаемся выдать
	return s.rRepo.Issue(ctx, userID, bookID)
}

func (s *rentalService) ReturnBook(ctx context.Context, userID, bookID int) error {
	return s.rRepo.Return(ctx, userID, bookID)
}
