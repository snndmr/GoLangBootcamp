package models

import (
	"errors"
	"strconv"
)

type Book struct {
	Id         int
	Name       string
	StockId    int
	ISBN       string
	PageCount  int
	Price      float32
	StockCount int
	Author     Author
	IsDeleted  bool
}

// NewBook constructor of Book
func NewBook(id, stockId, pageCount, stockCount int, name, ISBN string, price float32, author Author) *Book {
	return &Book{
		Id:         id,
		Name:       name,
		StockId:    stockId,
		ISBN:       ISBN,
		PageCount:  pageCount,
		Price:      price,
		StockCount: stockCount,
		Author:     author,
		IsDeleted:  false,
	}
}

type Deletable interface {
	Delete()
}

// Delete to delete book
func (book *Book) Delete() error {
	if book.IsDeleted {
		return errors.New("No book found with this id! ")
	}

	book.IsDeleted = true
	return nil
}

// Buy to buy books, error returns if stock is not enough
func (book *Book) Buy(buyCount int) error {
	if book.IsDeleted {
		return errors.New("No book found with this id! ")
	}

	if buyCount < 1 {
		return errors.New("You must enter the number of purchases as a positive integer. ")
	}

	if book.StockCount < buyCount {
		return errors.New("There is not enough stock. (Existing stock : " + strconv.Itoa(book.StockCount) + ")")
	}

	book.StockCount -= buyCount
	return nil
}
