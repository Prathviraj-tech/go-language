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

	minSalary := emp[0].salary
	for i := 1; i < n; i++ {
		if emp[i].salary < minSalary {
			minSalary = emp[i].salary
		}
	}

	fmt.Println("Employees with Minimum Salary:")
	for i := 0; i < n; i++ {
		if emp[i].salary == minSalary {
			fmt.Println(emp[i])
		}
	}
}