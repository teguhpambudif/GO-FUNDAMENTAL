package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	// berikut merupakan kombinasi slice dan struct, ditandai dengan "[]"
	var allStudents = []person{
		{name: "Wick", age: 23},
		{name: "Ethan", age: 23},
		{name: "Bourne", age: 23},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}
}
