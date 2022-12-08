package main

import "fmt"

func main() {
	// simple ways to looping array, using "for - range"
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for i, fruit := range fruits { // "fruit" to save per-element from fruits, index save on variable "i"
		fmt.Printf("element %d : %s\n", i, fruit)
	}
}
