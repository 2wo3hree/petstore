package facade

import (
	"context"
	"petstore/internal/models"
	"petstore/internal/service"
)

type Facade struct {
	library service.LibraryService
}

func NewFacade(lib service.LibraryService) *Facade {
	return &Facade{library: lib}
}

func (f *Facade) IssueBook(ctx context.Context, userID, bookID int) (models.Rental, error) {
	return f.library.IssueBook(ctx, userID, bookID)
}

func (f *Facade) ReturnBook(ctx context.Context, userID, bookID int) error {
	return f.library.ReturnBook(ctx, userID, bookID)
}

func (f *Facade) TopAuthors(ctx context.Context) ([]models.AuthorCount, error) {
	return f.library.GetTopAuthors(ctx)
}
