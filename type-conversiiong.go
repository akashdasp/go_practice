package main

import (
	"fmt"
	"math"
)

func main() {
	x, y := 1, 2
	f := math.Sqrt(float64(x*x + y*y))
	z := uint(f)
	fmt.Println(x, y, f, z)

}
