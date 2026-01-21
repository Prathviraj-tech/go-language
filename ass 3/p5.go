package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i <= 10; i++ {
		fmt.Println(i)

		delay := time.Duration(rand.Intn(251)) * time.Millisecond
		time.Sleep(delay)
	}
}
