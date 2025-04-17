package models

type User struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	RentedBooks []Book `json:"rented_books" db:"-"`
}

type ListUsersResponse struct {
	Total int    `json:"total"`
	Users []User `json:"users"`
}

type CreateUserRequest struct {
	Name string `json:"name" example:"Иван Иванов"`
}
