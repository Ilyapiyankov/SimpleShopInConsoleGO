package main

type Book struct {
	name  string
	price int
}

func MakeBook(name string, pr, count int) Book {
	return Book{name: name, price: pr}
}
