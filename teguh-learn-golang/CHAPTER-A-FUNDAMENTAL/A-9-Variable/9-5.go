package main

import "fmt"

func main() {
	// variable "_" will not be used aka dump variable
	_ = "belajar golang"
	_ = "golang itu mudah"
	name, _ := "john", "wick"

	fmt.Println(name)
}
