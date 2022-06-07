package main

import "fmt"

type student struct {
	name  string
	grade int
}

// variabel objek pointer
func main() {
	var s1 = student{name: "wick", grade: 2}

	var s2 *student = &s1 // the value of s2 is reference from s1
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 4, name :", s2.name) // accessing property without deferencing

	s2.name = "ethan" // change value directly using priginal value
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 4, name :", s2.name)
}
