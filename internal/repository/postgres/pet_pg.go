package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"petstore/internal/models"
	"petstore/internal/repository"
)

type petRepo struct {
	db *pgxpool.Pool
}

func NewPetRepo(db *pgxpool.Pool) repository.PetRepository {
	return &petRepo{db: db}
}

func (r *petRepo) Create(ctx context.Context, pet models.Pet) error {
	query := `INSERT INTO pets (id, name, status) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, query, pet.ID, pet.Name, pet.Status)
	return err
}

func (r *petRepo) GetByID(ctx context.Context, id int64) (models.Pet, error) {
	query := `SELECT id, name, status FROM pets WHERE id=$1`
	row := r.db.QueryRow(ctx, query, id)

	var pet models.Pet
	err := row.Scan(&pet.ID, &pet.Name, &pet.Status)
	return pet, err
}

func (r *petRepo) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM pets WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *petRepo) Update(ctx context.Context, pet models.Pet) error {
	query := `UPDATE pets SET name=$1, status=$2 WHERE id=$3`
	_, err := r.db.Exec(ctx, query, pet.Name, pet.Status, pet.ID)
	return err
}

func (r *petRepo) List(ctx context.Context, status string) ([]models.Pet, error) {
	query := `SELECT id, name, status FROM pets WHERE status=$1`
	rows, err := r.db.Query(ctx, query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pets []models.Pet
	for rows.Next() {
		var pet models.Pet
		if err := rows.Scan(&pet.ID, &pet.Name, &pet.Status); err != nil {
			return nil, err
		}
		pets = append(pets, pet)
	}
	return pets, nil
}
