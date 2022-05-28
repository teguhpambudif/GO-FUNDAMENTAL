package main

import "fmt"

func main() {
	/*
	   slice operation, slice is an reference per element from an array
	   ex:
	   var fruits = []string{"apple", "grape", "banana", "melon"}
	   var newFruits = fruits[0:2]
	*/

	var fruits = []string{"apple", "grape", "banana", "melon"}

	fmt.Println(fruits[0:2])
	fmt.Println(fruits[0:4])
	fmt.Println(fruits[0:0])
	fmt.Println(fruits[4:4])
	// fmt.Println(fruits[4:0]/[b:a]) this one will error, bcuz "a must be <= b"
	fmt.Println(fruits[:])
	fmt.Println(fruits[2:])
	fmt.Println(fruits[:2])

}
