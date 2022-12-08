package main

import "fmt"

func main() {
	var numberA int = 4
	var numberB *int = &numberA // "&" for referencing, and "*" for dereferencing

	fmt.Println("numberA (value) :", numberA)
	fmt.Println("numberA (address) :", &numberA)

	fmt.Println("numberB (value) :", *numberB)
	fmt.Println("numberB (address) :", numberB)
}
