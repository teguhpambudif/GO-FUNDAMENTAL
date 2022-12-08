/*
EMBEDED STRUCT
mechanism for embedding a struct as a property of another struct.
*/

package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	grade  int
	person // The person struct is embedded into the student struct, embeded struct is mutable
}

func main() {
	var s1 = student{}
	s1.name = "wick"
	s1.age = 21
	s1.grade = 2

	fmt.Println("name   :", s1.name)
	fmt.Println("age    :", s1.age)
	fmt.Println("age    :", s1.person.age)
	fmt.Println("grade  :", s1.grade)
}
