package main

import "fmt"

func main() {
	var firstName string = "Teguh"

	var lastName string
	lastName = "Pambudi"

	// Printf & Println, Printf didn't add new line
	fmt.Printf("halo Teguh Pambudi!\n")
	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println("halo", firstName, lastName+"!")
}
