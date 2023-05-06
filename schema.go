package main

import (
	"errors"
	"fmt"

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
	},
})
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
				"authorId": &graphql.ArgumentConfig{
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
			},
			Resolve: addAuthor,
		},
	},
})

var authorType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Author",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"dateOfBirth": &graphql.Field{
			Type: graphql.String,
		},
		"id": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var bookType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"numberOfChapters": &graphql.Field{
			Type: graphql.Int,
		},
		"numberOfPages": &graphql.Field{
			Type: graphql.Int,
		},
		"plot": &graphql.Field{
			Type: graphql.String,
		},
		"author": &graphql.Field{
			Type: authorType,
		},
	},
})

func getBookById(params graphql.ResolveParams) (interface{}, error) {
	id, ok := params.Args["id"].(int)
	if ok {
		// Find book by ID
		for _, book := range books {
			if book.ID == id {
				return book, nil
			}
		}
		return nil, errors.New("Book not found")
	}
	return nil, errors.New("invalid ID")
}
func getBooks(params graphql.ResolveParams) (interface{}, error) {
	return books, nil
}

func addAuthor(params graphql.ResolveParams) (interface{}, error) {
	name, _ := params.Args["name"].(string)
	datefBirth, _ := params.Args["dateOfBirth"].(string)
	newAuthor := Author{ID: len(authors) + 1, Name: name, DateOfBirth: datefBirth}
	authors = append(authors, newAuthor)
	return newAuthor, nil

}
func addBook(params graphql.ResolveParams) (interface{}, error) {
	authorId, ok := params.Args["authorId"].(int)
	var authorVal Author
	if ok {
		for _, author := range authors {
			if author.ID == authorId {
				authorVal = author
				title, _ := params.Args["title"].(string)
				numberOfChapters, _ := params.Args["numberOfChapters"].(int)
				numberOfPages, _ := params.Args["numberOfPages"].(int)

				newBook := Book{ID: len(books) + 1, Title: title, Author: &authorVal, NumberOfChapters: numberOfChapters, NumberOfPages: numberOfPages}
				books = append(books, newBook)
				fmt.Println(len(books))
				return newBook, nil
			}
		}
		return nil, errors.New("Book not found")

	}
	return nil, errors.New("invalid ID")

}

func getAuthors(params graphql.ResolveParams) (interface{}, error) {
	return authors, nil
}

func getAuthorById(params graphql.ResolveParams) (interface{}, error) {
	authorId, ok := params.Args["id"].(int)
	if ok {
		for _, author := range authors {
			if author.ID == authorId {
				return author, nil
			}
		}
		return nil, errors.New("Author not found")
	}
	return nil, errors.New("invalid ID")
}
