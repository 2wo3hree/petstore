package models

type Author struct {
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Books []Book `json:"books" db:"-"`
}

type AuthorCount struct {
	Author Author `json:"author"`
	Count  int64  `json:"count"`
}

type CreateAuthorRequest struct {
	Name string `json:"name"`
}
