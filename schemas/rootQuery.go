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
			Type:    graphql.NewList(bookType),
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
	},
})
