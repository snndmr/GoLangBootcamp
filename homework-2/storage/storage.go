package storage

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-2-snndmr/models"
	"strconv"
	"strings"
)

var books []models.Book
var authors []models.Author

// GetAuthors returns the authors
func GetAuthors() []models.Author {
	return authors
}

// InsertAuthor adds the author to the storage
func InsertAuthor(author *models.Author) {
	authors = append(authors, *author)
}

// GetBooks returns the books
func GetBooks() []models.Book {
	return books
}

// InsertBook adds the book to the storage
func InsertBook(book *models.Book) {
	books = append(books, *book)
}

// GetBookById returns the book according to the book id given as the argument
func GetBookById(id int) *models.Book {
	for index, book := range books {
		if book.Id == id {
			return &books[index]
		}
	}
	return nil
}

// SearchInBooks returns the ids of the books found by searching in the book name,
// ISBN, author name and stock id according to the substr given as an argument.
func SearchInBooks(substr string) []int {
	lowerSubstr := strings.ToLower(substr)

	var ids []int
	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Name), lowerSubstr) ||
			strings.Contains(strings.ToLower(book.ISBN), lowerSubstr) ||
			strings.Contains(strings.ToLower(book.Author.Name), lowerSubstr) ||
			strings.Contains(strings.ToLower(strconv.Itoa(book.StockId)), lowerSubstr) {
			ids = append(ids, book.Id)
		}
	}
	return ids
}
