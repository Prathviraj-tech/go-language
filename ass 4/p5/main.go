package main

import (
	"fmt"
	"rectarea/rectangle"
)

func main() {
	var l, w float64

	fmt.Print("Enter length: ")
	fmt.Scanln(&l)

	fmt.Print("Enter width: ")
	fmt.Scanln(&w)

	area := rectangle.Area(l, w)
	fmt.Println("Area of Rectangle =", area)
}
