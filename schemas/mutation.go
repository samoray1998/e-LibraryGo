package schemas

import (
	"github.com/graphql-go/graphql"
)

var mMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"addBook": &graphql.Field{
			Type: bookType,
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"numberOfChapters": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"numberOfPages": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"isbn": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"imageUrl": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"publishDate": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"authorId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"genreId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: addBook,
		},
		"addAuthor": &graphql.Field{
			Type: authorType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"dateOfBirth": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"boigraphy": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: addAuthor,
		},
		"addGenre": &graphql.Field{
			Type: genreType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: addGenre,
		},
		"regester": &graphql.Field{
			Type: userType,
			Args: graphql.FieldConfigArgument{
				"userName": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: regester,
		},
	},
})
