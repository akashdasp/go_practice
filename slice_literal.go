package main

import "fmt"

func main() {
	q := []int{2, 3, 4, 5, 6}
	fmt.Println(q)
	r := []bool{true, false, true, false}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, true},
		{4, false},
	}
	fmt.Println(len(s))
}
