package main

import "fmt"

func main() {
	var firstName string = "teguh"
	lastName := "pambudi" // type inference, no need type to be written

	fmt.Printf("halo %s %s!\n", firstName, lastName)
}
