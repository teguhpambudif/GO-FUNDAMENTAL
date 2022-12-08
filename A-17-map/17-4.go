package main

import "fmt"

func main() {
	// delete function on map, input map object and the key
	var chicken = map[string]int{"januari": 50, "februari": 40}

	fmt.Println(len(chicken))
	fmt.Println(chicken)

	delete(chicken, "januari")

	fmt.Println(len(chicken))
	fmt.Println(chicken)
}
