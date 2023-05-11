package schemas

import (
	"errors"
	"testGoGraph/models"
	"time"

	"github.com/graphql-go/graphql"
)

var reviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Review",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"user": &graphql.Field{
			Type: userType,
		},
		"book": &graphql.Field{
			Type: bookType,
		},
		"rating": &graphql.Field{
			Type: graphql.Float,
		},
		"textReview": &graphql.Field{
			Type: graphql.String,
		},
		"createdAt": &graphql.Field{
			Type: graphql.DateTime,
		},
		"updatedAt": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var reviews = []models.Review{
	{
		ID:         1,
		User:       &users[0],
		Book:       &books[0],
		Rating:     4.5,
		TextReview: "This book was really great!",
		CreatedAt:  time.Now().Add(-24 * time.Hour),
		UpdatedAt:  time.Now().Add(-12 * time.Hour),
	},
	{
		ID:         2,
		User:       &users[0],
		Book:       &books[0],
		Rating:     3.0,
		TextReview: "I thought this book was just okay.",
		CreatedAt:  time.Now().Add(-48 * time.Hour),
		UpdatedAt:  time.Now().Add(-24 * time.Hour),
	},
}

func getRviews(params graphql.ResolveParams) (interface{}, error) {
	return reviews, nil
}

func getBookReviews(params graphql.ResolveParams) (interface{}, error) {
	bookId, ok := params.Args["bookId"].(int)
	var bookReviews = []models.Review{}
	if ok {
		for _, review := range reviews {
			if review.Book.ID == bookId {
				bookReviews = append(bookReviews, review)
			}
		}
		if len(bookReviews) > 0 {
			return bookReviews, nil
		}
		return bookReviews, nil

	}

	return nil, errors.New("invalid passed id, try again")
}
