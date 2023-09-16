package main

import "fmt"

const pi = 3.14

func main() {
	const hello = "Hello"
	fmt.Println("Hi", hello)
	fmt.Println("this is ", pi, "day")
	const Truth = true
	fmt.Println("this is", Truth)
}
