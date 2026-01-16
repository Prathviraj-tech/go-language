package main

import "fmt"


func calculate(a int, b int) (int, int, int) {
	return a + b, a - b, a * b
}

func main() {
	var x, y int
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&x, &y)

	sum, diff, prod := calculate(x, y)

	fmt.Println("Sum =", sum)
	fmt.Println("Difference =", diff)
	fmt.Println("Product =", prod)
}
