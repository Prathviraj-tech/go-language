package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("data.txt")
	defer f.Close()

	b := make([]byte, 100)
	n, _ := f.Read(b)
	fmt.Println(string(b[:n]))
}
