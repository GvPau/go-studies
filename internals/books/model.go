package books

type Book struct {
	ID     string  `json:"id" validate:"required"`
	Title  string  `json:"title" validate:"required"`
	Author string  `json:"author" validate:"required"`
	Price  float64 `json:"price" validate:"required"`
}

var books = []Book{
	{ID: "1", Title: "Go Programming", Author: "John Doe", Price: 3.99},
	{ID: "2", Title: "Book 2", Author: "Author B", Price: 3.99},
	{ID: "3", Title: "Book 3", Author: "Author A", Price: 3.99},
	{ID: "4", Title: "Book 4", Author: "Author A", Price: 3.99},
}
