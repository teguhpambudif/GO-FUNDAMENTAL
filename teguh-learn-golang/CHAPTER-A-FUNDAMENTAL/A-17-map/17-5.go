package main

import "fmt"

func main() {
	// detect item on map with intended key
	var chicken = map[string]int{"januari": 50, "februari": 40}
	var value, isExist = chicken["mei"] // isExist as item accessed holder

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exist")
	}
}
