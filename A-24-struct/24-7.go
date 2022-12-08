/*
Anonymous struct is directly declared during object making
*/

package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var s1 = struct { // this is anonymous struct which have property "person(embeded)"", and "grade"
		person
		grade int
	}{} // declaration must be followed by initialization. "{}"
	s1.person = person{"wick", 21}
	s1.grade = 2

	fmt.Println("name    :", s1.person.name)
	fmt.Println("age     :", s1.person.age)
	fmt.Println("grade   :", s1.grade)

	// anonymous dengan pengisian properti
	var s2 = struct {
		person
		grade int
	}{
		person: person{"wick", 21},
		grade:  2,
	}

	fmt.Println("name   :", s2.person.name)
	fmt.Println("age    :", s2.person.age)
	fmt.Println("grade  :", s2.grade)
}
