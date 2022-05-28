package main

import "fmt"

func main() {
	/*
	   when len == cap, element from append become new reference on the last slice
	   when len < cap, element append will change the reference before append
	*/
	var fruits = []string{"apple", "grape", "banana"}
	var bFruits = fruits[0:2]

	fmt.Println(cap(bFruits))
	fmt.Println(len(bFruits))
	fmt.Println()

	fmt.Println(fruits)
	fmt.Println(bFruits)
	fmt.Println()

	// after this append, the last element from fruits will change
	var cFruits = append(bFruits, "papaya")

	fmt.Println(fruits)
	fmt.Println(bFruits)
	fmt.Println(cFruits)
}
