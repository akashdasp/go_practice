package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")
	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["p"] = "python"
	languages["go"] = "golang"
	fmt.Println(languages)
	fmt.Println(languages["p"])
	// delete(languages, "p")
	fmt.Println("After Deletting  one key from map", languages)
	for key, value := range languages {
		fmt.Printf("for key %v and value %v\n", key, value)
	}
}
