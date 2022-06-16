package main

import (
	"fmt"
	"strings"
)

type student struct {
	name  string
	grade int
}

// to declare method it is necessary to add an object variable declaration in between the func keyword and the function name. The "struct" used will be the owner of the method.
func (s student) sayHello() {
	fmt.Println("halo", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

func main() {
	var s1 = student{"john wick", 21}
	s1.sayHello()

	var name = s1.getNameAt(2)
	fmt.Println("nama panggilan :", name)
}
