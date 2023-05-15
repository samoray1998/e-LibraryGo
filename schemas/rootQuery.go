package schemas

import (
	"github.com/graphql-go/graphql"
)

var xRootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"book": &graphql.Field{
			Type: bookType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: getBookById,
		},
		"books": &graphql.Field{
			Type: graphql.NewList(bookType),
			Args: graphql.FieldConfigArgument{
				"token": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: getBooks,
		},
		"authors": &graphql.Field{
			Type:    graphql.NewList(authorType),
			Resolve: getAuthors,
		},
		"author": &graphql.Field{
			Type: authorType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: getAuthorById,
		},
		"genres": &graphql.Field{
			Type: graphql.NewList(
				genreType,
			),
			Resolve: getGenres,
		},
		"genre": &graphql.Field{
			Type: genreType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: getGenreById,
		},
		"users": &graphql.Field{
			Type: graphql.NewList(
				userType,
			),
			Resolve: getUsers,
		},

		"login": &graphql.Field{
			Type: authPayload,
			Args: graphql.FieldConfigArgument{
				"userName": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: logInUser,
		},
		"reviews": &graphql.Field{
			Type: graphql.NewList(
				reviewType,
			),
			Resolve: getRviews,
		},
		"borrowings": &graphql.Field{
			Type: graphql.NewList(
				borrowingType,
			),
			Resolve: getBorrowings,
		},
		"userBorrowings": &graphql.Field{
			Type: graphql.NewList(
				borrowingType,
			),
			Args: graphql.FieldConfigArgument{
				"userId": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: getUserBorrowings,
		},
		"bookReviews": &graphql.Field{
			Type: graphql.NewList(
				reviewType,
			),
			Args: graphql.FieldConfigArgument{
				"bookId": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: getBookReviews,
		},
		"notifications": &graphql.Field{
			Type: graphql.NewList(
				notificationType,
			),
			Resolve: getNotifications,
		},
	},
})
