package main

import "fmt"

func main() {
	var str1, str2 string

	
	fmt.Print("Enter first string: ")
	fmt.Scan(&str1)

	fmt.Print("Enter second string: ")
	fmt.Scan(&str2)

	
	if str1 == str2 {
		fmt.Println("Both strings are equal.")
	} else if str1 > str2 {
		fmt.Println("First string is greater than second string.")
	} else {
		fmt.Println("First string is less than second string.")
	}
}
