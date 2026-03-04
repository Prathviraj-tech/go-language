package main

import "fmt"


func sumOfSquares(num int, ch chan int) {
	sum := 0
	for num > 0 {
		d := num % 10
		sum += d * d
		num /= 10
	}
	ch <- sum
}


func sumOfCubes(num int, ch chan int) {
	sum := 0
	for num > 0 {
		d := num % 10
		sum += d * d * d
		num /= 10
	}
	ch <- sum
}

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	squareChan := make(chan int)
	cubeChan := make(chan int)

	go sumOfSquares(number, squareChan)
	go sumOfCubes(number, cubeChan)

	fmt.Println("Sum of squares:", <-squareChan)
	fmt.Println("Sum of cubes:", <-cubeChan)
}
