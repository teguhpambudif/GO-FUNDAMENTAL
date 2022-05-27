package main

import "fmt"

func main() {
	// for loop on array
	// the first is using itterator variable to access element based on index
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("element %d : %s\n", i, fruits[i])
	}
}
