package main
import "fmt"

func main() {
	var n int
	fmt.Print("Enter number of elements: ")
	fmt.Scan(&n)

	nums := make([]int, n)

	fmt.Println("Enter elements:")
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Println("Original slice:", nums)

	var newElement int
	fmt.Print("Enter element to append: ")
	fmt.Scan(&newElement)

	nums = append(nums, newElement)
	fmt.Println("After append:", nums)

	var index int
	fmt.Print("Enter index to remove: ")
	fmt.Scan(&index)

	if index >= 0 && index < len(nums) {
		nums = append(nums[:index], nums[index+1:]...)
		fmt.Println("After remove:", nums)
	} else {
		fmt.Println("Invalid index")
	}

	copySlice := make([]int, len(nums))
	copy(copySlice, nums)
	fmt.Println("Copied slice:", copySlice)
}