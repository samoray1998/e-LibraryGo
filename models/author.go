package models

type Author struct {
	Name        string `json:"name"`
	DateOfBirth string `json:"dateOfBirth"`
	ID          int    `json:"id"`
	Boigraphy   string `json:"boigraphy"`
}
