package models

import "time"

type Rental struct {
	ID           int        `json:"id" db:"id"`
	UserID       int        `json:"user_id" db:"user_id"`
	BookID       int        `json:"book_id" db:"book_id"`
	DateIssued   time.Time  `json:"date_issued" db:"date_issued"`
	DateReturned *time.Time `json:"date_returned,omitempty" db:"date_returned"`
	User         *User      `json:"user,omitempty" db:"-"`
	Book         *Book      `json:"book,omitempty" db:"-"`
}
