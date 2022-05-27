package main

import "fmt"

func main() {
	// array with vertical style
	var fruits [4]string

	fruits = [4]string{
		"apple",
		"grape",
		"banana",
		"melon", // "," still be used even it is the last element of array

	}
	fmt.Println(fruits)
}
