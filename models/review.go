package models

import "time"

type Review struct {
	ID         int       `json:"id"`
	User       *User     `json:"user"`
	Book       *Book     `json:"book"`
	Rating     float64   `json:"rating"`
	TextReview string    `json:"textReview"`
	CreatedAt  time.Time `json:"craetesAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
