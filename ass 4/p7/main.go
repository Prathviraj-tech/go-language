package main

import (
	"fmt"
	"calculatorApp/calculator"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b)
	fmt.Scan(&c)

	switch c {
	case 1:
		fmt.Println(calculator.Add(a, b))
	case 2:
		fmt.Println(calculator.Sub(a, b))
	case 3:
		fmt.Println(calculator.Mul(a, b))
	case 4:
		fmt.Println(calculator.Div(a, b))
	}
}
