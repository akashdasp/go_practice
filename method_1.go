package main

import "fmt"

func main() {
	fmt.Println("struct in golang")
	akash := User{"Akash", "akash.dasp@gmail.com", true, 12, 1}
	fmt.Println(akash)
	fmt.Println(akash.Status)
	fmt.Println(akash.Name)
	fmt.Println(akash.oneage)
	akash.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	oneage int
}

func (u User) GetStatus() {
	fmt.Println("User is Acitve:", u.Status)
}
