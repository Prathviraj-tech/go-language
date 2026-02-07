package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var filename string
	fmt.Print("Enter file name: ")
	fmt.Scanln(&filename)

	info, err := os.Stat(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(info.Name())
	fmt.Println(info.Size())
	fmt.Println(info.Mode())
	fmt.Println(info.ModTime().Format(time.RFC1123))
	fmt.Println(info.IsDir())
}
