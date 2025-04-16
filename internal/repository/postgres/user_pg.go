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

func (r *userRepo) Create(ctx context.Context, user models.User) error {
	query := `INSERT INTO users (id, username, first_name, last_name, email, password, phone, user_status)
			  VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := r.db.Exec(ctx, query, user.ID, user.Username, user.FirstName, user.LastName,
		user.Email, user.Password, user.Phone, user.UserStatus)
	return err
}

func (r *userRepo) GetByUsername(ctx context.Context, username string) (models.User, error) {
	query := `SELECT id, username, first_name, last_name, email, password, phone, user_status FROM users WHERE username=$1`
	row := r.db.QueryRow(ctx, query, username)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName,
		&user.Email, &user.Password, &user.Phone, &user.UserStatus)
	return user, err
}

func (r *userRepo) Delete(ctx context.Context, username string) error {
	query := `DELETE FROM users WHERE username=$1`
	_, err := r.db.Exec(ctx, query, username)
	return err
}

func (r *userRepo) Update(ctx context.Context, username string, updated models.User) error {
	query := `UPDATE users SET first_name = $1, last_name = $2, email = $3 WHERE username = $4`

	_, err := r.db.Exec(ctx, query, updated.FirstName, updated.LastName, updated.Email, username)
	return err
}
