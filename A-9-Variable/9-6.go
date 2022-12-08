package main

import "fmt"

func main() {
	name := new(string)

	fmt.Println(name)  // pointer to hexadecimanl
	fmt.Println(*name) // deferenced pointer

}
