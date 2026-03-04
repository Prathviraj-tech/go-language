package main
import "fmt"

type Student struct {
	rollNo  int
	name    string
	mark1   float64
	mark2   float64
	mark3   float64
	total   float64
	average float64
}

func main() {
	var n int
	fmt.Print("Enter number of students: ")
	fmt.Scan(&n)

	students := make([]Student, n)

	for i := 0; i < n; i++ {
		fmt.Println("Enter RollNo, Name, Mark1, Mark2, Mark3:")
		fmt.Scan(&students[i].rollNo, &students[i].name, &students[i].mark1, &students[i].mark2, &students[i].mark3)

		students[i].total = students[i].mark1 + students[i].mark2 + students[i].mark3
		students[i].average = students[i].total / 3
	}

	fmt.Println("Student Details:")
	for i := 0; i < n; i++ {
		fmt.Println(students[i])
	}
}