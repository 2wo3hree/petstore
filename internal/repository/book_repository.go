package repository

import (
	"context"
	"errors"
	"petstore/internal/models"
)

var ErrAlreadyIssued = errors.New("книга уже выдана")

type BookRepository interface {
	Create(ctx context.Context, b models.Book) (int64, error)
	GetByID(ctx context.Context, id int) (models.Book, error)
	List(ctx context.Context, limit, offset int) ([]models.Book, int, error)
	IssueBook(ctx context.Context, userID, bookID int) (models.Rental, error)
	ReturnBook(ctx context.Context, userID, bookID int) error
}
