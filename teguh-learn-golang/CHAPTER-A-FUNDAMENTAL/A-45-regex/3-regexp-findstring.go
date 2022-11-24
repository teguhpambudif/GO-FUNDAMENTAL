package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Digunakan untuk mencari string yang memenuhi kriteria regexp yang telah ditentukan.
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	var str = regex.FindString(text)
	fmt.Println(str)
	fmt.Printf("%T\n", str)
}
