package main

import "fmt"

func main() {
	// map with horizontal style
	var chicken1 = map[string]int{"januari": 50, "februari": 40}

	// map with vertical style
	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}

	var chicken3 = map[string]int{}
	var chicken4 = make(map[string]int)
	var chicken5 = *new(map[string]int) // with new it will make pointer, so we need to add "*" to see what it is

	fmt.Println(chicken1)
	fmt.Println(chicken2)
	fmt.Println(chicken3)
	fmt.Println(chicken4)
	fmt.Println(chicken5)

}
