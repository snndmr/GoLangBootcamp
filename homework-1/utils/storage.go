package utils

import "Library/models"

var books = []models.Book{
	{"0195153448", "Classical Mythology", "Mark P.O.Morford", 2002},
	{"0002005018", "Clara Callan", "Richard Bruce Wright", 2001},
	{"0060973129", "Decision in Normandy", "Carlo D'Este", 1991},
	{"0887841740", "The Middle Stories", "Sheila Heti", 2004},
	{"0425163091", "Chocolate Jesus", "Stephan Jaramillo", 1998},
	{"055321215X", "Pride and Prejudice", "Jane Austen", 1983},
	{"0553582909", "Icebound", "Dean R. Koontz", 2000},
	{"0441783589", "Starship Troopers", "Robert A. Heinlein", 1987},
	{"0446610399", "The Rescue", "Nicholas Sparks", 2001},
	{"0786863986", "A Monk Swimming", "Malachy McCourt", 1998},
}

// GetBook returns the book based on the index value given as an argument.
func GetBook(index int) models.Book {
	return books[index]
}

// GetBooks returns all books as an array.
func GetBooks() []models.Book {
	return books
}
