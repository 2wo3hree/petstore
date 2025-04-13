package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"repo/internal/model"
	"repo/internal/repository"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) repository.UserRepository {
	return &userRepo{db: db}
}

func (u *userRepo) Create(ctx context.Context, user model.User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2)`
	_, err := u.db.Exec(ctx, query, user.Name, user.Email)
	return err
}

func (u *userRepo) GetByID(ctx context.Context, id string) (model.User, error) {
	var user model.User
	query := `SELECT id, name, email FROM users WHERE id = $1 AND deleted_at IS NULL`
	err := u.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email)
	return user, err
}

func (u *userRepo) Update(ctx context.Context, user model.User) error {
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3 AND deleted_at IS NULL`
	_, err := u.db.Exec(ctx, query, user.Name, user.Email, user.ID)
	return err
}

func (u *userRepo) Delete(ctx context.Context, id string) error {
	query := `UPDATE users SET deleted_at = NOW() WHERE id = $1`
	_, err := u.db.Exec(ctx, query, id)
	return err
}

func (u *userRepo) List(ctx context.Context, limit, offset int) ([]model.User, int, error) {
	query := `SELECT id, name, email, deleted_at FROM users WHERE deleted_at IS NULL ORDER BY id LIMIT $1 OFFSET $2`
	rows, err := u.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.DeletedAt); err != nil {
			return nil, 0, err
		}
		users = append(users, user)
	}

	countQuery := `SELECT COUNT(*) FROM users WHERE deleted_at IS NULL`
	var count int
	err = u.db.QueryRow(ctx, countQuery).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	return users, count, nil
}
