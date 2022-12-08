package main

import (
	"fmt"
	"strings"
)

func main() {
	var secret interface{}

	secret = 2
	var number = secret.(int) * 10 // casting into int
	fmt.Println(secret, "multiplied by 10 is :", number)
	fmt.Println()

	secret = []string{"apple", "manggo", "banana"}
	var gruits = strings.Join(secret.([]string), ", ") // casting into []string
	fmt.Println(gruits, "is my favorite fruits")
}
