package library

import "library/book"

type Library struct {
	books map[book.Book]uint
}

func NewLibrary() *Library {
	return &Library{
		books: make(map[book.Book]uint, 10),
	}
}

//func (l *Library) is

func (l *Library) AddBook(name string, author string, amount uint) {

}
