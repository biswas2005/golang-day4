package functions

import "fmt"

func reference(a, b *int) int {
	*a = *a * 2
	return *a * *b

}

func Reference() {
	x := 5
	y := 10
	fmt.Printf("Before value of x=%d, y=%d\n", x, y)
	result := reference(&x, &y)
	fmt.Printf("Multiplication is: %d\n", result)
	fmt.Printf("After values of x=%d ,y=%d\n", x, y)
}
