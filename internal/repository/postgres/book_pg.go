package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"petstore/internal/models"
	"petstore/internal/repository"
)

type bookRepo struct {
	db *pgxpool.Pool
}

func NewBookRepo(pool *pgxpool.Pool) repository.BookRepository {
	return &bookRepo{db: pool}
}

func (r *bookRepo) Create(ctx context.Context, b models.Book) (int64, error) {
	const q = `
		INSERT INTO books (title, author_id)
		VALUES ($1, $2)
		RETURNING id
	`
	var id int64
	err := r.db.QueryRow(ctx, q, b.Title, b.AuthorID).Scan(&id)
	return id, err
}

func (r *bookRepo) GetByID(ctx context.Context, id int) (models.Book, error) {
	const q = `
		SELECT 
			b.id, b.title, b.author_id,
			a.id, a.name
		FROM books b
		JOIN authors a ON b.author_id = a.id
		WHERE b.id = $1
	`
	row := r.db.QueryRow(ctx, q, id)

	var b models.Book
	var a models.Author
	err := row.Scan(
		&b.ID, &b.Title, &b.AuthorID,
		&a.ID, &a.Name,
	)
	if err != nil {
		return models.Book{}, err
	}
	a.Books = nil // чтобы не пушить рекурсию
	b.Author = &a
	return b, nil
}

func (r *bookRepo) List(ctx context.Context, limit, offset int) ([]models.Book, int, error) {
	const dataQ = `
		SELECT 
			b.id, b.title, b.author_id,
			a.id, a.name
		FROM books b
		JOIN authors a ON b.author_id = a.id
		ORDER BY b.id
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.Query(ctx, dataQ, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	books := []models.Book{}
	for rows.Next() {
		var b models.Book
		var a models.Author
		if err := rows.Scan(
			&b.ID, &b.Title, &b.AuthorID,
			&a.ID, &a.Name,
		); err != nil {
			return nil, 0, err
		}
		a.Books = nil
		b.Author = &a
		books = append(books, b)
	}

	const countQ = `SELECT COUNT(*) FROM books`
	var total int
	if err := r.db.QueryRow(ctx, countQ).Scan(&total); err != nil {
		return nil, 0, err
	}

	return books, total, nil
}
