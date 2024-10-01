package main

import (
	"library/book"
	"library/library"
)

func main() {
	var library1 *library.Library = library.NewLibrary()
	library1.AddBook(*book.NewBook("Harry Potter", "J.K. Rowling"), 10)
	library1.AddBook(*book.NewBook("Little Women", "L.M. Alcott"), 8)
	library1.DisplayBooks()
}
