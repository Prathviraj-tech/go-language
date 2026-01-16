package main

import "fmt"

func main() {
	var n int
	var a, b, next int

	fmt.Print("Enter number of terms: ")
	_, err := fmt.Scan(&n)
	if err != nil || n <= 0 {
		fmt.Println("Please enter a valid number of terms greater than 0.")
		return
	}

	a = 0
	b = 1

	fmt.Println("Fibonacci Series:")

	for i := 1; i <= n; i++ {
		fmt.Print(a, " ")
		next = a + b
		a = b
		b = next
	}
	fmt.Println()
}
