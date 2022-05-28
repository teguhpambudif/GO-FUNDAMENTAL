package main

import "fmt"

func main() {
	/*
	   slice just like array and map, but we dont need to write the "element count"
	   ex:
	   var fruitsA = []string{"apple", "grape"}		this is slice
	   var fruitsB = [2]string{"banana", "melon"}		this is array
	   var fruitsC = [...]string{"papaya", "grape"}	this is another array
	*/

	var fruits = []string{"apple", "grape", "banana", "melon"}

	fmt.Println(fruits[0])
}
