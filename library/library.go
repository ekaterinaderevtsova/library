package library

import (
	"fmt"
	"library/book"
)

type Library struct {
	books map[book.Book]uint
}

func NewLibrary() *Library {
	return &Library{
		books: make(map[book.Book]uint, 10),
	}
}

//func (l *Library) is

func (l *Library) AddBook(book book.Book, num uint) {
	oldNum, exists := l.books[book]
	if exists {
		l.books[book] = oldNum + num
	} else {
		l.books[book] = num
	}
}

func (l *Library) RemoveAllCopiesBook(book book.Book) {
	_, exists := l.books[book]
	if exists {
		delete(l.books, book)
	} else {
		fmt.Println("No such book in the library.")
	}
}
