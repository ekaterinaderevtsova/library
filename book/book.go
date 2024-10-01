package book

type Book struct {
	title  string
	author string
}

func NewBook(title string, author string) *Book {
	return &Book{
		title:  title,
		author: author,
	}
}
