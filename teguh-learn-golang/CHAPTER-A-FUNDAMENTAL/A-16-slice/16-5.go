package main

import "fmt"

func main() {
	// cap function is to count maximum capacity of the slice
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))
	fmt.Println()

	var aFruits = fruits[0:3]
	fmt.Println(len(aFruits))
	fmt.Println(cap(aFruits))
	fmt.Println()

	var bFruits = fruits[1:4] // in this cap become 3 bcuz fruits[1] become bFruits[0]
	fmt.Println(len(bFruits))
	fmt.Println(cap(bFruits))
	fmt.Println()

}
