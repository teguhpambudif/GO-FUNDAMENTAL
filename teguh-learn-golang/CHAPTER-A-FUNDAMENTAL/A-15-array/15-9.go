package main

import "fmt"

func main() {
	// declaration array using keyword "make"
	var fruits = make([]string, 2)
	fruits[0] = "apple"
	fruits[1] = "manggo"

	fmt.Println(fruits)
}
