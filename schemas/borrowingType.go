package schemas

import (
	"errors"
	//"fmt"
	"testGoGraph/models"
	"time"

	"github.com/graphql-go/graphql"
)

var statusType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Status",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var borrowingType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Borrowing",
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
		"borrowingDate": &graphql.Field{
			Type: graphql.DateTime,
		},
		"returnDate": &graphql.Field{
			Type: graphql.DateTime,
		},
		"status": &graphql.Field{
			Type: statusType,
		},
	},
})

var borrowings = []models.Borrowing{
	{
		ID:           1,
		User:         &users[0],
		Book:         &books[0],
		BorronigDate: time.Date(2023, 05, 07, 11, 05, 51, 854880000, time.UTC),

		ReturnDate: time.Date(2023, 5, 17, 0, 0, 0, 0, time.UTC),
		Status:     &status[0],
	},
	{
		ID:           2,
		User:         &users[2],
		Book:         &books[2],
		BorronigDate: time.Date(2023, 05, 07, 11, 05, 51, 854880000, time.UTC),

		ReturnDate: time.Date(2023, 5, 17, 0, 0, 0, 0, time.UTC),
		Status:     &status[3],
	},
}

var status = []models.Status{
	{ID: 1, Name: "borrowed"},
	{ID: 2, Name: "returned"},
	{ID: 3, Name: "overdue"},
	{ID: 4, Name: "lost"},
	{ID: 5, Name: "damaged"}}

func getBorrowings(params graphql.ResolveParams) (interface{}, error) {
	return borrowings, nil
}

func getUserBorrowings(params graphql.ResolveParams) (interface{}, error) {
	userId, ok := params.Args["userId"].(int)

	var myBorrowings = []models.Borrowing{}
	if ok {

		for _, borrowing := range borrowings {
			if borrowing.User.ID == userId {
				myBorrowings = append(myBorrowings, borrowing)
				// return borrowing, nil
			}
		}
		if len(myBorrowings) > 0 {
			return myBorrowings, nil
		}
		return nil, errors.New("not found please try again")

	}
	return nil, errors.New("invalid passed id, try again")
}
