package main

import "fmt"

func main() {
	var secret interface{}

	secret = "ethan hunt"
	fmt.Println("data 1   : ", secret)

	secret = []string{"apple", "manggo", "banana"}
	fmt.Println("data 2   : ", secret)

	secret = 12.4
	fmt.Println("data 3   : ", secret)
}
