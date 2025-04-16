package service

import (
	"context"
	"petstore/internal/models"
	"petstore/internal/repository"
)

type PetService interface {
	Create(ctx context.Context, pet models.Pet) error
	GetByID(ctx context.Context, id int64) (models.Pet, error)
	Update(ctx context.Context, pet models.Pet) error
	Delete(ctx context.Context, id int64) error
	FindByStatus(ctx context.Context, status string) ([]models.Pet, error)
}

type petService struct {
	repo repository.PetRepository
}

func NewPetService(r repository.PetRepository) PetService {
	return &petService{repo: r}
}

func (s *petService) Create(ctx context.Context, pet models.Pet) error {
	// Здесь может быть валидация
	return s.repo.Create(ctx, pet)
}

func (s *petService) GetByID(ctx context.Context, id int64) (models.Pet, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *petService) Update(ctx context.Context, pet models.Pet) error {
	return s.repo.Update(ctx, pet)
}

func (s *petService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *petService) FindByStatus(ctx context.Context, status string) ([]models.Pet, error) {
	return s.repo.List(ctx, status)
}
