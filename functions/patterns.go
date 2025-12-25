package functions

import "fmt"

func Pattern() {
	var n int
	fmt.Print("Enter any number: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print("1")

		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()

	}
}
