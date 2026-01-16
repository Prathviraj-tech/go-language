package main

import "fmt"


func calculate(a int, b int) (int, int, int, float64) {
	add := a + b
	sub := a - b
	mul := a * b

	var div float64
	if b != 0 {
		div = float64(a) / float64(b)
	}

	return add, sub, mul, div
}

func main() {
	var x, y int
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&x, &y)

	add, sub, mul, div := calculate(x, y)

	fmt.Println("Addition =", add)
	fmt.Println("Subtraction =", sub)
	fmt.Println("Multiplication =", mul)

	if y != 0 {
		fmt.Println("Division =", div)
	} else {
		fmt.Println("Division not possible (division by zero).")
	}
}
