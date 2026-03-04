package main

import "fmt"

func main() {
	var num int

	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if (num >= 10 && num <= 99) || (num <= -10 && num >= -99) {
		fmt.Println("The number is a two-digit number.")
	} else {
		fmt.Println("The number is not a two-digit number.")
	}
}
