package main

import "fmt"

func main() {

	evenSum := 0
	oddSum := 0

	for i := 1; i <= 100; i++ {

		if i%2 == 0 {
			evenSum += i
		} else {
			oddSum += i
		}
	}

	fmt.Printf("Sum of all even numbers between 1 and 100 is: %d\n", evenSum)
	fmt.Printf("Sum of all odd numbers between 1 and 100 is: %d\n", oddSum)
}
