package main

import (
	"fmt"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-2-snndmr/models"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-2-snndmr/storage"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

	// To statically initialize data
	storage.InsertAuthor(models.NewAuthor(1, "Mark P.O.Mor ford"))
	storage.InsertAuthor(models.NewAuthor(2, "Richard Bruce Wright"))
	storage.InsertAuthor(models.NewAuthor(3, "Carlo D'Est"))
	storage.InsertAuthor(models.NewAuthor(4, "Sheila Hetti"))
	storage.InsertAuthor(models.NewAuthor(5, "Stephan Jaramillo"))

	storage.InsertBook(models.NewBook(1, 10, rand.Intn(1e3), rand.Intn(10), "Classical Mythology", "0195153448", 1e2*rand.Float32(), storage.GetAuthors()[0]))
	storage.InsertBook(models.NewBook(2, 20, rand.Intn(1e3), rand.Intn(10), "Clara Calla", "0002005018", 1e2*rand.Float32(), storage.GetAuthors()[1]))
	storage.InsertBook(models.NewBook(3, 30, rand.Intn(1e3), rand.Intn(10), "Decision in Normandy", "0060973129", 1e2*rand.Float32(), storage.GetAuthors()[2]))
	storage.InsertBook(models.NewBook(4, 40, rand.Intn(1e3), rand.Intn(10), "The Middle Stories", "0887841740", 1e2*rand.Float32(), storage.GetAuthors()[3]))
	storage.InsertBook(models.NewBook(5, 50, rand.Intn(1e3), rand.Intn(10), "Chocolate Jesus", "0425163091", 1e2*rand.Float32(), storage.GetAuthors()[4]))
}

func main() {
	args := os.Args

	// To check if arguments are entered
	if len(args) == 1 {
		fmt.Printf("Parameters you can use in the %s application:\n%-10s%s\n%-10s%s", filepath.Base(args[0]), "search", "To search books by name.", "list", "To list the books.")
		return
	}

	// To operate by parameter.
	switch strings.ToLower(args[1]) {
	case "buy":
		// To check if the book id is given as an argument
		if len(args) == 2 {
			fmt.Println("You must enter book id. (Ex: go run main.go buy 5 12)")
			return
		}

		// To check the validity of the book id
		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("You must enter valid book id. (Ex: go run main.go buy 5 12)")
			return
		}

		// To check if the number of books is given as an argument
		if len(args) == 3 {
			fmt.Println("You must enter the number of books you want to buy. (Ex: go run main.go buy 5 12)")
			return
		}

		// To check the validity of the number of books to be purchased
		count, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Println("You must enter valid book count. (Ex: go run main.go buy 5 12)")
			return
		}

		book := storage.GetBookById(id)
		if book == nil {
			fmt.Println("No book found with this id.")
			return
		}

		err = book.Buy(count)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Book purchased successfully!")
	case "delete":
		// To check if the book id is given as an argument
		if len(args) == 2 {
			fmt.Println("You must enter book id. (Ex: go run main.go delete 5)")
			return
		}

		// To check the validity of the book id
		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Printf("You must enter valid book id. (Ex: go run main.go delete 5)")
			return
		}

		book := storage.GetBookById(id)
		if book == nil {
			fmt.Println("No book found with this id.")
			return
		}

		err = book.Delete()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Book deleted successfully!")
	case "search":
		// To check if the substr is given as an argument
		if len(args) == 2 {
			fmt.Println("You must enter a value. (Ex: go run main.go search Chocolate Jesus)")
			return
		}

		// To concatenate and search disjointed strings
		bookSubstr := strings.Join(args[2:], " ")
		bookIds := storage.SearchInBooks(bookSubstr)

		if len(bookIds) == 0 {
			fmt.Printf("%s not found.", bookSubstr)
			return
		}

		// To print undeleted and searched books to the screen.
		fmt.Printf("%-5s%-30s%-15s%-30s%-15s%-10s%-10s%-10s", "#", "Name", "Stock ID", "Author Name", "ISBN", "Pages", "Price", "Stock")
		for _, id := range bookIds {
			book := storage.GetBookById(id)
			if !book.IsDeleted {
				fmt.Printf("\n%-5d%-30s%-15d%-30s%-15s%-10d%-10.2f%-10d", book.Id, book.Name, book.StockId, book.Author.Name, book.ISBN, book.PageCount, book.Price, book.StockCount)
			}
		}
	case "list":
		// To check if arguments are entered
		if len(args) > 2 {
			fmt.Println("List has no additional arguments (Ex: go run main.go list)")
			return
		}

		// To print undeleted books to the screen.
		fmt.Printf("%-5s%-30s%-15s%-30s%-15s%-10s%-10s%-10s", "#", "Name", "Stock ID", "Author Name", "ISBN", "Pages", "Price", "Stock")
		for _, book := range storage.GetBooks() {
			if !book.IsDeleted {
				fmt.Printf("\n%-5d%-30s%-15d%-30s%-15s%-10d%-10.2f%-10d", book.Id, book.Name, book.StockId, book.Author.Name, book.ISBN, book.PageCount, book.Price, book.StockCount)
			}
		}
	default:
		fmt.Printf("%s unknown parameter!", args[1])
		return
	}
}
