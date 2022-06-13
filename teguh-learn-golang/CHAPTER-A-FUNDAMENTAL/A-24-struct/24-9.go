package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var allStudents = []struct { // anonynous struct
		person
		grade int
	}{
		{person: person{"wick", 21}, grade: 2},
		{person: person{"ethan", 22}, grade: 2},
		{person: person{"bond", 21}, grade: 2},
	}

	for _, student := range allStudents {
		fmt.Println(student)
	}
}
