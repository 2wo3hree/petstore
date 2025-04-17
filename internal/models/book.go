package models

type Book struct {
	ID       int     `json:"id" db:"id"`
	Title    string  `json:"title" db:"title"`
	AuthorID int     `json:"author_id" db:"author_id"`
	Author   *Author `json:"author,omitempty" db:"-"`
}

type CreateBookRequest struct {
	Title    string `json:"title"`
	AuthorID int    `json:"author_id"`
}

type ListBooksResponse struct {
	Total int    `json:"total"`
	Books []Book `json:"books"`
}
