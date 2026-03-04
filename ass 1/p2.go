package main

import "fmt"

func main() {
	var number int
	fmt.Println("Enter a number: ")
	fmt.Scanln(&number)

	if number%2 == 0 {
		fmt.Println("The number is even.")
	} else {
		fmt.Println("The number is odd.")
	}
}
