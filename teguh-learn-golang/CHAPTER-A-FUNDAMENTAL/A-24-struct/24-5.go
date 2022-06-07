/*
Embedded Struct With Same Property Name
bcuz property name is the same, so accessing the property must be done explicitly or clearly
*/

package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	person // embeded struct
	age    int
	grade  int
}

func main() {
	var s1 = student{}
	s1.name = "wick"
	s1.age = 21
	s1.person.age = 22

	fmt.Println(s1.name)
	fmt.Println(s1.age)
	fmt.Println(s1.person.age) // accessing property "age" from struct "person" via "object struct student"
}
