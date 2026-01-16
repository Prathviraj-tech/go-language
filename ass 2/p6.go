package main

import "fmt"


func addSub(a int, b int) (int, int) {
	return a + b, a - b
}

func main() {
	var x, y int
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&x, &y)

	sum, diff := addSub(x, y)

	fmt.Println("Addition =", sum)
	fmt.Println("Subtraction =", diff)
}
