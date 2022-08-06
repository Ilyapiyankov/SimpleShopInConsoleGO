package main

type User struct {
	balance     int
	BoughtBooks map[Book]int
}

func (us *User) InStock(book Book) bool {
	for i, _ := range us.BoughtBooks {
		if i == book {
			return true
		}
	}
	return false
}

func CreatUser(balance int) User {
	return User{balance: balance, BoughtBooks: map[Book]int{}}
}

func (us *User) addBook(shop *Shop, bookName string, number int) (bool, string) {

	if shop.IsInStock(bookName) {

		book := shop.GetBook(bookName)

		if us.balance >= book.price*number && shop.BooksInStock[book] >= number {
			flg := false
			for i, _ := range us.BoughtBooks {
				if i == book {
					us.BoughtBooks[book] += number
					us.balance -= book.price * number
					shop.BooksInStock[book] -= number
					flg = true
					break
				}
			}
			if !flg {
				us.balance -= book.price * number
				shop.BooksInStock[book] -= number
				us.BoughtBooks[shop.GetBook(bookName)] = number
			}
			return true, ""
		} else {
			if us.balance < book.price*number {
				return false, "Not enough money to buy this book"
			}
		}
	} else {
		return false, "We don't have this book in stock"
	}
	return false, "Как сэда дошло? Проверь User.addBook - там ошибка!"
}
