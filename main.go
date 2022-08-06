package main

import (
	"fmt"
	"strings"
)

func main() {
	var balance int
	ShopMain := createShop(nil)

	fmt.Print("Enter your balance (in rubles): ")
	fmt.Scanln(&balance)
	fmt.Println()

	if balance == 0 || balance < ShopMain.minPrice {
		fmt.Println("Maybe your input is incorrect or not enough money, so your balance is 1000 now")
		balance = 1000
	}

	Usr := CreatUser(balance)
	//initializing shop and user

	fmt.Printf("Your balance: %v, you have no books\n", Usr.balance)

	//saying hello
	fmt.Println(strings.Repeat("-", 50) + "Hello from simple Console Shop on golang by Ilya Piyankov" + strings.Repeat("-", 50))

	fmt.Println(strings.Repeat("-", 55) + "Books in stock, their prices, how much in stock" + strings.Repeat("-", 55))

	for i, o := range ShopMain.BooksInStock {
		fmt.Printf("Book's name: '%v', price: %v, in stock: %v \n", i.name, i.price, o)
	}
	/*fmt.Println("buying book#1 in 2 copies")

	fmt.Println(Usr.addBook(&ShopMain, "Book#1", 2))

	fmt.Println(Usr.balance, ShopMain.BooksInStock[ShopMain.GetBook("Book#1")], Usr.BoughtBooks)*/
	LoopFunc(&Usr, &ShopMain)
}
