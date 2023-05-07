package schemas

import (
	"errors"
	"fmt"
	"testGoGraph/models"

	"github.com/graphql-go/graphql"
)

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
		"isbn": &graphql.Field{
			Type: graphql.String,
		},
		"publishDate": &graphql.Field{
			Type: graphql.String,
		},
		"imageUrl": &graphql.Field{
			Type: graphql.String,
		},
		"genre": &graphql.Field{
			Type: genreType,
		},
		// "genre":&graphql.Field{
		// 	Type:
		// }
	},
})

var books = []models.Book{
	{
		ID:               1,
		Title:            "To Kill a Mockingbird",
		NumberOfChapters: 31,
		NumberOfPages:    281,
		Author:           &authors[0],
		Plot:             "The story takes place in the fictional town of Maycomb, Alabama during the Great Depression, and follows the story of Scout Finch, her brother Jem, and their father Atticus, a lawyer who defends a black man accused of raping a white woman.",
		Isbn:             "irut94579",
		ImageURl:         "https://images-na.ssl-images-amazon.com/images/I/61PYe5p5l5L._SX329_BO1,204,203,200_.jpg",
		PublishDate:      "2020-08-13",
		Genre:            &genres[0],
	},
	{
		ID:               2,
		Title:            "The Great Gatsby",
		NumberOfChapters: 9,
		NumberOfPages:    180,
		Author:           &authors[1],
		Plot:             "The novel takes place in the summer of 1922 in the fictional town of West Egg on Long Island. It tells the story of a mysterious millionaire, Jay Gatsby, and his obsession with the beautiful former debutante Daisy Buchanan.",
		Isbn:             "ooiretj93489",
		ImageURl:         "https://images-na.ssl-images-amazon.com/images/I/61PYe5p5l5L._SX329_BO1,204,203,200_.jpg",
		PublishDate:      "1998-07-18",
		Genre:            &genres[2],
	},
	{
		ID:               3,
		Title:            "Pride and Prejudice",
		NumberOfChapters: 61,
		NumberOfPages:    279,
		Author:           &authors[2],
		Plot:             "The novel follows the story of Elizabeth Bennet, the second of five daughters of a country gentleman. When wealthy Mr. Bingley and his friend Mr. Darcy arrive in town, Mrs. Bennet sees a chance for her daughters to marry into wealth.",
		Isbn:             "9847590",
		ImageURl:         "https://images-na.ssl-images-amazon.com/images/I/61PYe5p5l5L._SX329_BO1,204,203,200_.jpg",
		PublishDate:      "2023-08-13",
		Genre:            &genres[1],
	},
}

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

func addBook(params graphql.ResolveParams) (interface{}, error) {
	authorId, ok := params.Args["authorId"].(int)
	var authorVal models.Author
	if ok {
		for _, author := range authors {
			if author.ID == authorId {
				authorVal = author
				title, _ := params.Args["title"].(string)
				numberOfChapters, _ := params.Args["numberOfChapters"].(int)
				numberOfPages, _ := params.Args["numberOfPages"].(int)
				isbn, _ := params.Args["isbn"].(string)
				imgSrc, _ := params.Args["imageUrl"].(string)
				publishDate, _ := params.Args["publishDate"].(string)

				newBook := models.Book{ID: len(books) + 1, Title: title, Author: &authorVal, NumberOfChapters: numberOfChapters, NumberOfPages: numberOfPages, Isbn: isbn, ImageURl: imgSrc, PublishDate: publishDate}
				books = append(books, newBook)
				fmt.Println(len(books))
				return newBook, nil
			}
		}
		return nil, errors.New("Book not found")

	}
	return nil, errors.New("invalid ID")

}
