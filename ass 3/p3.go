package main

import "fmt"


func evenReceiver(evenCh chan int) {
	fmt.Println("Even Numbers:")
	for num := range evenCh {
		fmt.Println(num)
	}
}


func oddReceiver(oddCh chan int) {
	fmt.Println("Odd Numbers:")
	for num := range oddCh {
		fmt.Println(num)
	}
}

func main() {
	var n int
	fmt.Print("Enter number of elements: ")
	fmt.Scan(&n)

	numbers := make([]int, n)

	fmt.Println("Enter elements:")
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	evenCh := make(chan int)
	oddCh := make(chan int)

	go evenReceiver(evenCh)
	go oddReceiver(oddCh)

	for _, num := range numbers {
		if num%2 == 0 {
			evenCh <- num
		} else {
			oddCh <- num
		}
	}

	close(evenCh)
	close(oddCh)
}
