package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"petstore/internal/models"
	"petstore/internal/repository"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) repository.UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, user models.User) (int64, error) {
	const q = `
    INSERT INTO users (name)
    VALUES ($1)
    RETURNING id
    `
	var id int64
	err := r.db.QueryRow(ctx, q, user.Name).Scan(&id)
	return id, err
}

func (r *userRepo) GetByID(ctx context.Context, id int) (models.User, error) {
	const sql = `SELECT id, name FROM users WHERE id = $1`
	var u models.User
	err := r.db.QueryRow(ctx, sql, id).Scan(&u.ID, &u.Name)
	return u, err
}

func (r *userRepo) List(ctx context.Context, limit, offset int) ([]models.User, int, error) {
	const dataQ = `
      SELECT id, name
      FROM users
      ORDER BY id
      LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(ctx, dataQ, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return nil, 0, err
		}
		users = append(users, u)
	}

	const countQ = `
      SELECT COUNT(*) 
      FROM users`

	var total int
	if err := r.db.QueryRow(ctx, countQ).Scan(&total); err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *userRepo) GetWithRentals(ctx context.Context, id int) (models.User, error) {
	// Сначала базовый пользователь
	u, err := r.GetByID(ctx, id)
	if err != nil {
		return u, err
	}

	// Теперь достаём все активные аренды (не возвращённые книги),
	// вместе с данными о книге и авторе.
	const sql = `
    SELECT
      b.id, b.title, b.author_id,
      a.id, a.name,
      r.id, r.user_id, r.book_id, r.date_issued, r.date_returned
    FROM rentals r
    JOIN books b   ON r.book_id   = b.id
    JOIN authors a ON b.author_id = a.id
    WHERE r.user_id = $1 AND r.date_returned IS NULL
    `
	rows, err := r.db.Query(ctx, sql, id)
	if err != nil {
		return u, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			book   models.Book
			author models.Author
			rent   models.Rental
		)
		// Сканим: book поля, author поля, rental поля
		if err := rows.Scan(
			&book.ID, &book.Title, &book.AuthorID,
			&author.ID, &author.Name,
			&rent.ID, &rent.UserID, &rent.BookID, &rent.DateIssued, &rent.DateReturned,
		); err != nil {
			return u, err
		}
		// Связываем книгу с автором, и добавляем книгу в список у пользователя
		author.Books = nil    // у книги нам не нужен её список
		book.Author = &author // вложенный автор
		rent.Book = &book     // вложенная книга
		rent.User = nil       // в rental поле User не используем
		u.RentedBooks = append(u.RentedBooks, book)
	}
	return u, rows.Err()
}
