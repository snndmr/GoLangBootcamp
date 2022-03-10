package main

import (
	"Library/helper"
	"Library/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		helper.ListCommands(args[0])
		return
	}

	switch strings.ToLower(args[1]) {
	case "search":
		if len(args) == 2 {
			fmt.Println("\n You must enter a value. (Ex: go run main.go search Chocolate Jesus)")
			return
		}

		// To combine sequential search arguments and search.
		bookSubstr := strings.Join(args[2:], " ")
		bookIndexes := utils.SearchInBooks(bookSubstr)

		if len(bookIndexes) == 0 {
			fmt.Printf("\n %s not found.", bookSubstr)
			fmt.Println()
			return
		}

		helper.ListSearchedBooks(bookIndexes)
	case "list":
		helper.ListAllBooks()
	default:
		fmt.Printf("\n %s unknown parameter!", args[1])
		helper.ListCommands(args[0])
		return
	}
}
