package functions

import "fmt"

func multiply(a, b int) int {
	a = a * 2
	return a * b

}
func Multiply() {
	x := 5
	y := 10
	fmt.Printf("Before x=%d, y=%d\n", x, y)
	result := multiply(x, y)
	fmt.Printf("Multiplication : %d\n", result)
	fmt.Printf("After x=%d, y=%d\n", x, y)
}
