package main

type Book struct {
	ID               int     `json:"id"`
	Title            string  `json:"title"`
	NumberOfChapters int     `json:"numberOfChapters"`
	NumberOfPages    int     `json:"numberOfPages"`
	Author           *Author `json:"author"`
	Plot             string  `json:"plot"`
}

var books = []Book{
	{
		ID:               1,
		Title:            "To Kill a Mockingbird",
		NumberOfChapters: 31,
		NumberOfPages:    281,
		Author:           &authors[0],
		// Author: &Author{
		// 	Name:        "Harper Lee",
		// 	DateOfBirth: "April 28, 1926",
		// 	ID:          1,
		// },
		Plot: "The story takes place in the fictional town of Maycomb, Alabama during the Great Depression, and follows the story of Scout Finch, her brother Jem, and their father Atticus, a lawyer who defends a black man accused of raping a white woman.",
	},
	{
		ID:               2,
		Title:            "The Great Gatsby",
		NumberOfChapters: 9,
		NumberOfPages:    180,
		Author:           &authors[1],
		Plot:             "The novel takes place in the summer of 1922 in the fictional town of West Egg on Long Island. It tells the story of a mysterious millionaire, Jay Gatsby, and his obsession with the beautiful former debutante Daisy Buchanan.",
	},
	{
		ID:               3,
		Title:            "Pride and Prejudice",
		NumberOfChapters: 61,
		NumberOfPages:    279,
		Author:           &authors[2],
		Plot:             "The novel follows the story of Elizabeth Bennet, the second of five daughters of a country gentleman. When wealthy Mr. Bingley and his friend Mr. Darcy arrive in town, Mrs. Bennet sees a chance for her daughters to marry into wealth.",
	},
}
