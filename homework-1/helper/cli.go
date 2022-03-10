package helper

import (
	"Library/utils"
	"fmt"
	"path/filepath"
)

// ListCommands informs the command line screen about commands that can be used.
func ListCommands(projectName string) {
	fmt.Printf("\n Parameters you can use in the %s application:\n %-10s%s\n %-10s%s", filepath.Base(projectName),
		"search", "To search books by name.",
		"list", "To list the books.")
	fmt.Println()
}

// ListAllBooks prints all books by formatting them on the command line.
func ListAllBooks() {
	fmt.Printf("\n%5s\t%-15s%-30s%-30s%s", "#", "ISBN", "Title", "Author", "Publication Year")
	for index, book := range utils.GetBooks() {
		fmt.Printf("\n%5d\t%-15s%-30s%-30s%d", index+1, book.ISBN, book.Title, book.Author, book.PublicationYear)
	}
	fmt.Println()
}

// ListSearchedBooks prints searched books by formatting them on the command line.
func ListSearchedBooks(indexes []int) {
	fmt.Printf("\n%5s\t%-15s%-30s%-30s%s", "#", "ISBN", "Title", "Author", "Publication Year")
	for index, index_ := range indexes {
		book := utils.GetBook(index_)
		fmt.Printf("\n%5d\t%-15s%-30s%-30s%d", index+1, book.ISBN, book.Title, book.Author, book.PublicationYear)
	}
	fmt.Println()
}
