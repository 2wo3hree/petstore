package model

type User struct {
	ID        string  `json:"id" db:"id"`
	Name      string  `json:"name" db:"name"`
	Email     string  `json:"email" db:"email"`
	DeletedAt *string `json:"deleted_at,omitempty" db:"deleted_at"`
}

type ListUsersResponse struct {
	Total int    `json:"total"`
	Users []User `json:"users"`
}
