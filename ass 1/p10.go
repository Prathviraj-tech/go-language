package main

import "fmt"


func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	return fact
}


func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var n int

	
	fmt.Print("Enter number of rows: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		
		for space := 1; space <= n-i; space++ {
			fmt.Print(" ")
		}

		
		for j := 0; j <= i; j++ {
			fmt.Printf("%d ", combination(i, j))
		}
		fmt.Println()
	}
}
