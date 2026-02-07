package main

import (
	"fmt"
	"square/sq"
)

func main() {
	var side float64

	fmt.Print("Enter side of square: ")
	fmt.Scanln(&side)

	fmt.Println("Area of Square =", sq.Area(side))
}
