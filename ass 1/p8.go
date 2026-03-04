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
	_, err := fmt.Scan(&x)
	if err != nil {
		fmt.Println("Invalid input for the first number.")
		return
	}

	fmt.Print("Enter second number: ")
	_, err = fmt.Scan(&y)
	if err != nil {
		fmt.Println("Invalid input for the second number.")
		return
	}

	
	fmt.Println("\nBefore swapping:")
	fmt.Println("x =", x, "y =", y)

	
	swap(&x, &y)

	
	fmt.Println("\nAfter swapping:")
	fmt.Println("x =", x, "y =", y)
}
