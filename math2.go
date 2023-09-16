package main

import "fmt"

func sum(x, y int) int {
	return x + y
}
func result_sum() {
	fmt.Println(sum(1, 3))
}
func main() {
	// Call your other functions here if needed
	result_sum()
}
