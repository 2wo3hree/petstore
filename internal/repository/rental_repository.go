package repository

import (
	"context"
	"errors"
	"petstore/internal/models"
)

var ErrAlreadyIssued = errors.New("книга уже выдана")

type RentalRepository interface {
	IssueBook(ctx context.Context, userID, bookID int) (models.Rental, error)
	ReturnBook(ctx context.Context, userID, bookID int) error
}
