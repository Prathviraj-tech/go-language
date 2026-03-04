package main
import "fmt"

type Book struct {
	BookID int
	Title  string
	Author string
	Price  float64
}

func main() {
	var n int
	fmt.Print("Enter number of books: ")
	fmt.Scan(&n)

	books := make([]Book, n)

	for i := 0; i < n; i++ {
		fmt.Println("Enter BookID, Title, Author, Price:")
		fmt.Scan(&books[i].BookID, &books[i].Title, &books[i].Author, &books[i].Price)
	}

	for i := 0; i < n; i++ {
		fmt.Println(books[i])
	}
}