package utils

import "strings"

// SearchInBooks searches for the substr value in book titles. It returns indexes of matches.
func SearchInBooks(substr string) []int {
	var indexes []int
	for index, book := range GetBooks() {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(substr)) {
			indexes = append(indexes, index)
		}
	}
	return indexes
}
