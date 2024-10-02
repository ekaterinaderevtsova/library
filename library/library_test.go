package library

import (
	"github.com/stretchr/testify/assert"
	"library/book"
	"testing"
)

type inputData struct {
	Book book.Book
	Num  uint
}

func TestLibrary_AddBook(t *testing.T) {

	t.Run("test 1", func(t *testing.T) {
		var library *Library = NewLibrary()
		var inputBook book.Book = *book.NewBook("Harry Potter", "J.K. Rowling")

		input := inputData{
			Book: inputBook,
			Num:  10,
		}

		library.AddBook(input.Book, input.Num)

		val, ok := library.books[input.Book]
		if ok && val == input.Num {
			assert.True(t, true)
		} else {
			assert.True(t, false)
		}
	})

	t.Run("test 2", func(t *testing.T) {
		var library *Library = NewLibrary()
		var inputBook book.Book = *book.NewBook("Harry Potter", "J.K. Rowling")

		library.AddBook(inputBook, 10)
		library.AddBook(inputBook, 5)

		val, ok := library.books[inputBook]
		if ok && val == 15 {
			assert.True(t, true)
		} else {
			assert.True(t, false)
		}
	})

	t.Run("test 3", func(t *testing.T) {
		var library *Library = NewLibrary()
		var inputBook1 book.Book = *book.NewBook("1984", "George Orwell")
		var inputBook2 book.Book = *book.NewBook("Little Women", "L.M. Alcott")
		input1 := inputData{
			Book: inputBook1,
			Num:  10,
		}

		input2 := inputData{
			Book: inputBook2,
			Num:  5,
		}

		library.AddBook(input1.Book, input1.Num)
		library.AddBook(input2.Book, input2.Num)

		val1, ok1 := library.books[inputBook1]
		val2, ok2 := library.books[inputBook2]
		if ok1 && ok2 && val1 == input1.Num && val2 == input2.Num {
			assert.True(t, true)
		} else {
			assert.True(t, false)
		}
	})
}

func TestLibrary_RemoveBook(t *testing.T) {
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

func TestLibrary_RemoveBookByTitle(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		var library *Library = NewLibrary()
		var inputBook book.Book = *book.NewBook("Harry Potter", "J.K. Rowling")

		library.AddBook(inputBook, 10)
		library.RemoveBookByTitle("Harry Potter")

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
		library.RemoveBookByTitle("Harry Potter")

		if len(library.books) == 1 {
			assert.True(t, true)
		} else {
			assert.True(t, false)
		}
	})

	t.Run("test3", func(t *testing.T) {
		var library *Library = NewLibrary()
		var inputBook1 book.Book = *book.NewBook("Harry Potter", "J.K. Rowling")
		var inputBook2 book.Book = *book.NewBook("Little Women", "L.M. Alcott")

		library.AddBook(inputBook1, 10)
		library.AddBook(inputBook2, 12)
		library.RemoveBookByTitle("Harry")

		if len(library.books) == 2 {
			assert.True(t, true)
		} else {
			assert.True(t, false)
		}
	})
}

func TestLibrary_BorrowBook(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		var library *Library = NewLibrary()
		var inputBook book.Book = *book.NewBook("Harry Potter", "J.K. Rowling")
		var n uint
		n = 10

		library.AddBook(inputBook, n)
		library.BorrowBook(inputBook)

		if library.books[inputBook] == n-1 {
			assert.True(t, true)
		} else {
			assert.True(t, false)
		}
	})

	t.Run("test2", func(t *testing.T) {
		var library *Library = NewLibrary()
		var inputBook1 book.Book = *book.NewBook("Harry Potter", "J.K. Rowling")
		var inputBook2 book.Book = *book.NewBook("Little Women", "L.M. Alcott")
		var n uint
		n = 10

		library.AddBook(inputBook1, n)
		library.AddBook(inputBook2, n)
		library.BorrowBook(inputBook1)

		if library.books[inputBook1] == n-1 {
			assert.True(t, true)
		} else {
			assert.True(t, false)
		}
	})

	t.Run("test3", func(t *testing.T) {
		var library *Library = NewLibrary()
		var inputBook1 book.Book = *book.NewBook("Harry Potter", "J.K. Rowling")
		var inputBook2 book.Book = *book.NewBook("Little Women", "L.M. Alcott")
		var n uint
		n = 10

		library.AddBook(inputBook1, n)
		library.AddBook(inputBook2, n)
		library.BorrowBook(*book.NewBook("Harry", "J.K. Rowling"))

		if library.books[inputBook1] == n {
			assert.True(t, true)
		} else {
			assert.True(t, false)
		}
	})
}
