package main

import "fmt"

func main() {
	// copy function is to copy element from src(2nd param), to dst(1st param). base on len dst
	// dst := make([]string, 3)
	// src := []string{"watermelon", "pinnaple", "apple", "orange"}
	// n := copy(dst, src) // function copy will return number

	// fmt.Println(dst)
	// fmt.Println(src)
	// fmt.Println(n)
	// fmt.Println()

	// ------
	// example below when variable src < or = dst

	dst := []string{"potato", "potato", "potato"}
	src := []string{"watermelon", "pinnaple"}
	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)
}
