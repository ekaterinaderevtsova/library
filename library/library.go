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

func (l *Library) AddBook(book book.Book, num uint) {
	oldNum, exists := l.books[book]
	if exists {
		l.books[book] = oldNum + num
	} else {
		l.books[book] = num
	}
}
