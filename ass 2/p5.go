package main

import "fmt"


func printTable(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	printTable(num)
}
