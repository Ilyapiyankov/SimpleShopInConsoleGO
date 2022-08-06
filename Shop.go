package main

import "fmt"

type Shop struct {
	BooksInStock map[Book]int
	minPrice     int
}

func createShop(Var map[Book]int) Shop {

	defaultStock := make(map[Book]int)
	if Var != nil {
		defaultStock = Var
	} else {
		for i := 1; i < 10; i++ {
			defaultStock[Book{price: i * 100, name: fmt.Sprintf("Book#%v", i)}] = 10 - i
		}
	}

	min := 0

	for i, _ := range defaultStock {
		min = i.price
		break
	}

	for i, _ := range defaultStock {
		if i.price < min {
			min = i.price
		}
	}

	return Shop{defaultStock, min}
}

func (receiver *Shop) IsInStock(bookName string) bool {

	for book, _ := range receiver.BooksInStock {
		if book.name == bookName {
			return true
		}
	}
	return false
}

func (receiver *Shop) GetBook(nm string) Book {
	v := Book{}
	for i, _ := range receiver.BooksInStock {
		if i.name == nm {
			v = i
			break
		}
	}
	return v
}
