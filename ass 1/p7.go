package main

import "fmt"

type Book struct {
	bookID int
	title  string
	author string
	price  float64
}

func main() {
	var n int

	fmt.Print("Enter number of books: ")
	_, err := fmt.Scan(&n)
	if err != nil || n <= 0 {
		fmt.Println("Please enter a valid number of books greater than 0.")
		return
	}

	books := make([]Book, n)

	for i := 0; i < n; i++ {
		fmt.Println("\nEnter details of Book", i+1)

		fmt.Print("Book ID: ")
		_, err := fmt.Scan(&books[i].bookID)
		if err != nil {
			fmt.Println("Invalid input for Book ID")
			return
		}

		fmt.Print("Title: ")
		fmt.Scanln(&books[i].title)

		fmt.Print("Author: ")
		fmt.Scanln(&books[i].author)

		fmt.Print("Price: ")
		_, err = fmt.Scan(&books[i].price)
		if err != nil {
			fmt.Println("Invalid input for Price")
			return
		}
	}

	fmt.Println("\nBook Details:")
	for i := 0; i < n; i++ {
		fmt.Println("\nBook", i+1)
		fmt.Println("Book ID :", books[i].bookID)
		fmt.Println("Title   :", books[i].title)
		fmt.Println("Author  :", books[i].author)

		fmt.Printf("Price   : %.2f\n", books[i].price)
	}
}
