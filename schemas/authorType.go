package schemas

import (
	"errors"
	"testGoGraph/models"

	"github.com/graphql-go/graphql"
)

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
		"boigraphy": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// Authores data

var authors = []models.Author{
	{
		Name:        "Jane Austen",
		DateOfBirth: "December 16, 1775",
		ID:          3,
		Boigraphy:   "Jane Austen was an English novelist known primarily for her six major novels, which interpret, critique and comment upon the British landed gentry at the end of the 18th century.",
	},
	{
		Name:        "F. Scott Fitzgerald",
		DateOfBirth: "September 24, 1896",
		ID:          2,
		Boigraphy:   "Jane Austen was an English novelist known primarily for her six major novels, which interpret, critique and comment upon the British landed gentry at the end of the 18th century.",
	},
	{
		Name:        "Jane test",
		DateOfBirth: "December 16, 1775",
		ID:          1,
		Boigraphy:   "Jane Austen was an English novelist known primarily for her six major novels, which interpret, critique and comment upon the British landed gentry at the end of the 18th century.",
	},
}

func addAuthor(params graphql.ResolveParams) (interface{}, error) {
	name, _ := params.Args["name"].(string)
	datefBirth, _ := params.Args["dateOfBirth"].(string)
	boi, _ := params.Args["boigraphy"].(string)
	newAuthor := models.Author{ID: len(authors) + 1, Name: name, DateOfBirth: datefBirth, Boigraphy: boi}
	authors = append(authors, newAuthor)
	return newAuthor, nil

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
