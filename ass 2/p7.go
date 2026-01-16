package main

import "fmt"

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	var x, y int

	fmt.Print("Enter first number: ")
	fmt.Scan(&x)
	fmt.Print("Enter second number: ")
	fmt.Scan(&y)

	fmt.Println("Before swapping: x =", x, "y =", y)
	swap(&x, &y)
	fmt.Println("After swapping:  x =", x, "y =", y)
}
