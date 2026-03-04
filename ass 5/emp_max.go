package main
import "fmt"

type Employee struct {
	eno    int
	ename  string
	salary float64
}

func main() {
	var n int
	fmt.Print("Enter number of employees: ")
	fmt.Scan(&n)

	emp := make([]Employee, n)

	for i := 0; i < n; i++ {
		fmt.Println("Enter Employee No, Name, Salary:")
		fmt.Scan(&emp[i].eno, &emp[i].ename, &emp[i].salary)
	}

	maxSalary := emp[0].salary
	for i := 1; i < n; i++ {
		if emp[i].salary > maxSalary {
			maxSalary = emp[i].salary
		}
	}

	fmt.Println("Employees with Maximum Salary:")
	for i := 0; i < n; i++ {
		if emp[i].salary == maxSalary {
			fmt.Println(emp[i])
		}
	}
}