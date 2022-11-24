package main

import "fmt"

func main() {
	// accessing slice element with 3 index
	var fruits = []string{"apple", "grape", "banana"}
	var aFruits = fruits[0:2]
	var bFruits = fruits[0:2:2] // 3rd index in this is the capacity

	fmt.Println(fruits)
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))
	fmt.Println()

	fmt.Println(aFruits)
	fmt.Println(len(aFruits))
	fmt.Println(cap(aFruits))
	fmt.Println()

	fmt.Println(bFruits)
	fmt.Println(len(bFruits))
	fmt.Println(cap(bFruits))
}
