package main

import (
	. "belajar-golang-level-akses/library"
	f "fmt" // using alias when import package
)

func main() {
	var s1 = Student{"ethan", 21} // bcoz we import the package using . notation, we dont need to write package name we use the struct
	f.Println("name", s1.Name)
	f.Println("grade", s1.Grade)
}
