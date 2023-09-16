package main

import "fmt"

var c, python, java bool
var i, j int = 1, 2

// short Hand declearation

func main() {
	var i int
	k := 3
	a, b := true, 2
	var c, js, rust = true, false, "no!"
	fmt.Println(i, c, python, java)
	fmt.Println(c, js, rust)
	fmt.Println(a, b, k)
}
