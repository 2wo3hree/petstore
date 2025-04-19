package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"petstore/internal/models"
	"petstore/internal/repository"
)

type rentalRepo struct {
	db *pgxpool.Pool
}

func NewRentalRepo(db *pgxpool.Pool) repository.RentalRepository {
	return &rentalRepo{db: db}
}

func (r *rentalRepo) IssueBook(ctx context.Context, userID, bookID int) (models.Rental, error) {
	const q = `
		INSERT INTO rentals (user_id, book_id)
		VALUES ($1, $2)
		RETURNING id, user_id, book_id, date_issued, date_returned
	`
	var rent models.Rental
	err := r.db.QueryRow(ctx, q, userID, bookID).
		Scan(&rent.ID, &rent.UserID, &rent.BookID, &rent.DateIssued, &rent.DateReturned)
	if err != nil {
		// если уникальный индекс нарушен
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return models.Rental{}, repository.ErrAlreadyIssued
		}

	}
	return rent, nil
}

func (r *rentalRepo) ReturnBook(ctx context.Context, userID, bookID int) error {
	const q = `
		UPDATE rentals
		SET date_returned = NOW()
		WHERE user_id = $1 AND book_id = $2 AND date_returned IS NULL
	`
	ct, err := r.db.Exec(ctx, q, userID, bookID)
	if err != nil {
		return err
	}
	if ct.RowsAffected() == 0 {
		return errors.New("никто не брал эту книгу")
	}
	return nil
}
