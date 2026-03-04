package main
import "fmt"

func main() {
	var rows, cols int

	fmt.Print("Enter number of rows: ")
	fmt.Scan(&rows)

	fmt.Print("Enter number of columns: ")
	fmt.Scan(&cols)

	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
	}

	fmt.Println("Enter matrix elements:")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	fmt.Println("Transpose Matrix:")
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			fmt.Print(matrix[j][i], " ")
		}
		fmt.Println()
	}
}