package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("sample.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	file.WriteString("\nThis line is appended using Go program.")
	fmt.Println("Done")
}
