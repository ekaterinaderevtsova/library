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

func (l *Library) RemoveBook(book book.Book) {
	_, exists := l.books[book]
	if exists {
		delete(l.books, book)
	} else {
		fmt.Println("No such book in the library.") //change to error
	}
}

func (l *Library) RemoveBookByTitle(title string) {
	for bookKey, _ := range l.books {
		if bookKey.GetTitle() == title {
			delete(l.books, bookKey)
			return
		}
	}
	fmt.Println("No such book in the library.") //change to error
}

func (l *Library) BorrowBook(book book.Book) {
	num, exists := l.books[book]
	if exists {
		l.books[book] = num - 1
	} else {
		fmt.Println("No such book in the library.") //change to error
	}
}

func (l *Library) BorrowBookByTitle(title string) {
	for bookKey, num := range l.books {
		if bookKey.GetTitle() == title {
			l.books[bookKey] = num - 1
			return
		}
	}
	fmt.Println("No such book in the library.") //change to error
}

func (l *Library) DisplayBooks() {
	for bookToPrint, num := range l.books {
		fmt.Println(bookToPrint.GetTitle(), ", ", bookToPrint.GetAuthor(), ", ", num, " copies")
	}
}
