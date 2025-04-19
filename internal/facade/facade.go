package facade

import (
	"context"
	"petstore/internal/models"
	"petstore/internal/service"
)

type Facade interface {
	Issue(ctx context.Context, userID, bookID int) (models.Rental, error)
	Return(ctx context.Context, userID, bookID int) error
	GetTopAuthors(ctx context.Context) ([]models.AuthorCount, error)
}

type FacadeImp struct {
	lib *service.LibrarySuperService
}

func NewFacadeImp(lib *service.LibrarySuperService) *FacadeImp {
	return &FacadeImp{lib: lib}
}

func (f *FacadeImp) Issue(ctx context.Context, userID, bookID int) (models.Rental, error) {
	return f.lib.Issue(ctx, userID, bookID)
}

func (f *FacadeImp) Return(ctx context.Context, userID, bookID int) error {
	return f.lib.Return(ctx, userID, bookID)
}

func (f *FacadeImp) GetTopAuthors(ctx context.Context) ([]models.AuthorCount, error) {
	return f.lib.GetTopAuthors(ctx)
}
