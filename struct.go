package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Printf("%t", v)
	fmt.Println("")
	v.X = 4
	v.Y = 5
	fmt.Println(v.X * v.Y)
}
