package main

import "fmt"


func isPalindrome(num int) bool {
	original := num
	reverse := 0

	for num > 0 {
		remainder := num % 10
		reverse = reverse*10 + remainder
		num = num / 10
	}

	return original == reverse
}

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	if isPalindrome(n) {
		fmt.Println("The number is a palindrome.")
	} else {
		fmt.Println("The number is not a palindrome.")
	}
}
