/*
Filling the sub-struct property values ​​can be done by directly entering the object variables that are printed from the same struct.
*/

package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	person
	age   int
	grade int
}

func main() {
	var p1 = person{name: "wick", age: 21}
	var s1 = student{person: p1, grade: 2}

	fmt.Println("name    :", s1.name)
	fmt.Println("age     :", s1.person.age) // in the tutorial fmt.Println("age   :", s1.age), but return "0"
	fmt.Println("grade   :", s1.grade)
}
