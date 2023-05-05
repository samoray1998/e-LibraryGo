package main

type Author struct {
	Name        string `json:"name"`
	DateOfBirth string `json:"dateOfBirth"`
	ID          int    `json:"id"`
}

// Authores data

var authors = []Author{
	{
		Name:        "Jane Austen",
		DateOfBirth: "December 16, 1775",
		ID:          3,
	},
	{
		Name:        "F. Scott Fitzgerald",
		DateOfBirth: "September 24, 1896",
		ID:          2,
	},
	{
		Name:        "Jane test",
		DateOfBirth: "December 16, 1775",
		ID:          1,
	},
}
