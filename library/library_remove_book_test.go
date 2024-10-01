package library

import (
	"github.com/stretchr/testify/assert"
	"library/book"
	"testing"
)

func TestLibrary_RemoveBook2Book(t *testing.T) {

	t.Run("test1", func(t *testing.T) {
		var library *Library = NewLibrary()
		var inputBook book.Book = *book.NewBook("Harry Potter", "J.K. Rowling")

		library.AddBook(inputBook, 10)
		library.RemoveBook(inputBook)

		if len(library.books) == 0 {
			assert.True(t, true)
		} else {
			assert.True(t, false)
		}
	})

	t.Run("test2", func(t *testing.T) {
		var library *Library = NewLibrary()
		var inputBook1 book.Book = *book.NewBook("Harry Potter", "J.K. Rowling")
		var inputBook2 book.Book = *book.NewBook("Little Women", "L.M. Alcott")

		library.AddBook(inputBook1, 10)
		library.AddBook(inputBook2, 12)
		library.RemoveBook(inputBook1)

		if len(library.books) == 1 {
			assert.True(t, true)
		} else {
			assert.True(t, false)
		}
	})
}
