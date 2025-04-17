package postgres

import (
	"context"
	"petstore/internal/models"
	"petstore/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type authorRepo struct {
	db *pgxpool.Pool
}

func NewAuthorRepo(pool *pgxpool.Pool) repository.AuthorRepository {
	return &authorRepo{db: pool}
}

func (r *authorRepo) Create(ctx context.Context, a models.Author) (int64, error) {
	const q = `INSERT INTO authors (name) VALUES ($1) RETURNING id`
	var id int64
	err := r.db.QueryRow(ctx, q, a.Name).Scan(&id)
	return id, err
}

func (r *authorRepo) GetByID(ctx context.Context, id int64) (models.Author, error) {
	const q = `SELECT id, name FROM authors WHERE id = $1`
	var a models.Author
	err := r.db.QueryRow(ctx, q, id).Scan(&a.ID, &a.Name)
	return a, err
}

func (r *authorRepo) List(ctx context.Context, limit, offset int) ([]models.Author, int, error) {
	const qData = `SELECT id, name FROM authors ORDER BY id LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(ctx, qData, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	authorsMap := make(map[int]*models.Author)
	var list []*models.Author
	for rows.Next() {
		var a models.Author
		if err := rows.Scan(&a.ID, &a.Name); err != nil {
			return nil, 0, err
		}
		a.Books = []models.Book{}
		authorsMap[a.ID] = &a
		list = append(list, &a)
	}

	// Подгружаем книги
	const qBooks = `SELECT id, title, author_id FROM books WHERE author_id = ANY($1)`
	ids := make([]int, 0, len(authorsMap))
	for id := range authorsMap {
		ids = append(ids, id)
	}
	bookRows, err := r.db.Query(ctx, qBooks, ids)
	if err != nil {
		return nil, 0, err
	}
	defer bookRows.Close()

	for bookRows.Next() {
		var b models.Book
		if err := bookRows.Scan(&b.ID, &b.Title, &b.AuthorID); err != nil {
			return nil, 0, err
		}
		if author, ok := authorsMap[b.AuthorID]; ok {
			author.Books = append(author.Books, b)
		}
	}

	const qCount = `SELECT COUNT(*) FROM authors`
	var total int
	if err := r.db.QueryRow(ctx, qCount).Scan(&total); err != nil {
		return nil, 0, err
	}

	// Преобразуем []*Author → []Author
	result := make([]models.Author, 0, len(list))
	for _, a := range list {
		result = append(result, *a)
	}
	return result, total, nil
}

func (r *authorRepo) TopAuthors(ctx context.Context) ([]models.AuthorCount, error) {
	const q = `
		SELECT 
  		a.id,
  		a.name,
  		COUNT(r.id) AS cnt
		FROM authors a
		JOIN books b ON b.author_id = a.id
		LEFT JOIN rentals r ON r.book_id = b.id
		GROUP BY a.id, a.name
		ORDER BY cnt DESC
		LIMIT 10
	`
	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []models.AuthorCount
	for rows.Next() {
		var ac models.AuthorCount
		if err := rows.Scan(&ac.Author.ID, &ac.Author.Name, &ac.Count); err != nil {
			return nil, err
		}
		out = append(out, ac)
	}
	return out, nil
}
