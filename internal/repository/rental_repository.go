package repository

import (
	"context"
	"errors"
	"petstore/internal/models"
)

var ErrAlreadyIssued = errors.New("книга уже выдана")

type RentalRepository interface {
	Issue(ctx context.Context, userID, bookID int) (models.Rental, error)
	Return(ctx context.Context, userID, bookID int) error
}
