package main

import (
	"fmt"
	"rectarea/rectangle"
)

func main() {
	var l, w float64
	fmt.Scanln(&l)
	fmt.Scanln(&w)
	fmt.Println(rectangle.Area(l, w))
}
