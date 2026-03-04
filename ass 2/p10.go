package main

import "fmt"


func changeValue(a int) {
	a = 100
	fmt.Println("Value inside function =", a)
}

func main() {
	x := 50

	fmt.Println("Value before function call =", x)
	changeValue(x)
	fmt.Println("Value after function call =", x)
}
