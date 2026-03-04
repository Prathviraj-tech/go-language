package main

import "fmt"

func main() {
	
	ch := make(chan int, 5)

	
	ch <- 10
	ch <- 20
	ch <- 30

	fmt.Println("Channel Capacity:", cap(ch))
	fmt.Println("Channel Length before reading:", len(ch))

	
	fmt.Println("Values read:")
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("Channel Length after reading:", len(ch))
}
