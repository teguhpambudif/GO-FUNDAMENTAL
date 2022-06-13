package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var student struct { // printed an object from anonymous struct then stored in a variable named student.
		person
		grade int
	} // we can also add an initialization for the anonymous struct right here like "{ 12, }"
	student.person = person{"wick", 21}
	student.grade = 2

	fmt.Println(student.name)
	fmt.Println(student.age)
	fmt.Println(student.grade)
}
