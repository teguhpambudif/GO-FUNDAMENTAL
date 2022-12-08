package main

import "fmt"

func main() {
	// if we only need elements data without index, wecan save index on "_" variable
	/* if we only need index, we can use 1 variable after keyword "for"
	example :
	for i, _ := range fruits {}
	or
	for i := range fruits {}

	*/
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	for _, fruit := range fruits {
		fmt.Printf("nama buah : %s\n", fruit)
	}
}
