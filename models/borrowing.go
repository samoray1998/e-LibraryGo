package models

import "time"

type Borrowing struct {
	ID           int       `json:"id"`
	User         *User     `json:"user"`
	Book         *Book     `json:"book"`
	BorronigDate time.Time `json:"borrowingDate"`
	ReturnDate   time.Time `json:"returnDate"`
	Status       *Status   `json:"status"`
}

type Status struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
