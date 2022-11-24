package main

import "fmt"

func main() {
	// this array have 4 capacity, starting from index 0, if you add more than 4 it will be an error
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])
}
