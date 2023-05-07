package models

type Book struct {
	ID               int     `json:"id"`
	Title            string  `json:"title"`
	NumberOfChapters int     `json:"numberOfChapters"`
	NumberOfPages    int     `json:"numberOfPages"`
	Author           *Author `json:"author"`
	Plot             string  `json:"plot"`
	Isbn             string  `json:"isbn"`
	PublishDate      string  `json:"publishDate"`
	ImageURl         string  `json:"imageUrl"`
}
