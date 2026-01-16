package main

import "fmt"

func main() {
	
	var studentName string
	var rollNo int
	var division string
	var collegeName string

	
	fmt.Print("Enter Student Name: ")
	fmt.Scanln(&studentName)

	
	fmt.Print("Enter Roll Number: ")
	fmt.Scanln(&rollNo)

	
	fmt.Print("Enter Division: ")
	fmt.Scanln(&division)

	
	fmt.Print("Enter College Name: ")
	fmt.Scanln(&collegeName)

	
	fmt.Println("\n--- Student Details ---")
	fmt.Printf("Name: %s\n", studentName)
	fmt.Printf("Roll Number: %d\n", rollNo)
	fmt.Printf("Division: %s\n", division)
	fmt.Printf("College Name: %s\n", collegeName)
}
