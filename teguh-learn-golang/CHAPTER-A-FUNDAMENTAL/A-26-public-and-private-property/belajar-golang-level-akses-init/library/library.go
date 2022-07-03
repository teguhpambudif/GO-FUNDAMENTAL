package library

import "fmt"

var Student = struct {
	Name  string
	Grade int
}{}

// this function will be executed first
func init() {
	Student.Name = "John Wick"
	Student.Grade = 2

	fmt.Println("--> library/library.go imported")
}
