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

func (book *Book) GetTitle() string {
	return book.title
}

func (book *Book) GetAuthor() string {
	return book.author
}
