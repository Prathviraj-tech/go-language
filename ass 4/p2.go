package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("sample.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	file.WriteString("Hello Students!\nThis file is created using Go language.")
	fmt.Println("File created")
}
