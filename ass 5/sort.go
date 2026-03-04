package main
import "fmt"

func main() {
	var n int
	fmt.Print("Enter number of elements: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	fmt.Println("Enter elements:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println("Sorted Array:")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
}