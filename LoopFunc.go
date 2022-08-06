package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoopFunc(user *User, shop *Shop) {

	commands := []string{
		"show balance",
		"show books in stock",
		"show my books",
		"_exit_",
		"add balance <value>",
		"buy book <nameofbook> <howmany>",
	}

	fmt.Println("Commands to use:\n")

	for o, i := range commands {
		fmt.Println(fmt.Sprintf("%v) %v\n", o, i))
	}
	fmt.Println("\n--------------------------------------\nexample of usage: buy book Book#1 2\n--------------------------------------")
	for command := enterCommand(); command != commands[3]; command = enterCommand() {
		if command == commands[0] {
			fmt.Println(user.balance)
			continue
		}
		if command == commands[1] {
			fmt.Println("------------------------------------Books in stock--------------------------------------")
			for i, o := range shop.BooksInStock {
				fmt.Printf("Book's name: '%v', price: %v, in stock: %v \n", i.name, i.price, o)
			}
			continue
		}
		if command == commands[2] {
			fmt.Println("--------------------------------------Your books----------------------------------------")
			if len(user.BoughtBooks) == 0 {
				fmt.Println("NO BOOKS!")
				continue
			}
			for i, o := range user.BoughtBooks {
				fmt.Printf("Book's name: '%v', price: %v, you have: %v \n", i.name, i.price, o)
			}
			continue
		} else {
			cmd := strings.Split(command, " ")

			cond1 := len(cmd) == 3 && cmd[0] == "add" && cmd[1] == "balance"
			cond2 := len(cmd) == 4 && cmd[0] == "buy" && cmd[1] == "book"

			switch {
			case cond1:
				ToAdd, err := strconv.Atoi(cmd[2])
				if err != nil {
					fmt.Println("Incorrect command input")
					break
				}
				user.balance += ToAdd
				fmt.Println("Your balance:", user.balance)
				break
			case cond2:
				BookName := cmd[2]
				HowMany, err := strconv.Atoi(cmd[3])
				if err != nil {
					fmt.Println("Incorrect command input")
					break
				}
				ok, str := user.addBook(shop, BookName, HowMany)
				if ok {
					fmt.Println("-----------------------------------success-----------------------------------")
					fmt.Println("your balance is", user.balance)
				} else {
					fmt.Println(str)
				}
				break
			default:
				fmt.Println("Incorrect command input")
			}
			continue
		}
	}
	fmt.Println("----END----")
}

func enterCommand() string {

	fmt.Print("Enter your command: ")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}
