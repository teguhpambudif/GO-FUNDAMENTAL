package main

import (
	"fmt"
	"strings"
)

// Variadic parameters can be combined with regular parameters, but variadic param must be positioned at the end
func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}

func main() {
	// yourHobbies("wick", "coding", "gaming")

	// if the second and so on parameters to be filled with data from the slice, then use a "..."
	var hobbies = []string{"sleeping", "eating"}
	yourHobbies("wick", hobbies...)
}
