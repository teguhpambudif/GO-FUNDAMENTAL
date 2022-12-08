package main

import "fmt"

func main() {
	// nested loop
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println() // Println without parameter is meant to add new line, just like Print("\n")
	}
}
