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

func (l *Library) RemoveBook(book book.Book) error {
	_, exists := l.books[book]
	if exists {
		delete(l.books, book)
		return nil
	} else {
		return fmt.Errorf("book not found in the library")
	}
}

func (l *Library) RemoveBookByTitle(title string) error {
	for bookKey, _ := range l.books {
		if bookKey.GetTitle() == title {
			delete(l.books, bookKey)
			return nil
		}
	}
	return fmt.Errorf("book with such title not found in the library")
}

func (l *Library) BorrowBook(book book.Book) error {
	num, exists := l.books[book]
	if exists {
		l.books[book] = num - 1
		return nil
	} else {
		return fmt.Errorf("book not found in the library")
	}
}

func (l *Library) BorrowBookByTitle(title string) error {
	for bookKey, num := range l.books {
		if bookKey.GetTitle() == title {
			l.books[bookKey] = num - 1
			return nil
		}
	}
	return fmt.Errorf("book with such title not found in the library")
}

func (l *Library) DisplayBooks() {
	for bookToPrint, num := range l.books {
		fmt.Println(bookToPrint.GetTitle(), ", ", bookToPrint.GetAuthor(), ", ", num, " copies")
	}
}
