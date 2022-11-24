package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Method ini digunakan untuk mendeteksi apakah string memenuhi sebuah pola regexp.
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	var isMatch = regex.MatchString(text)
	fmt.Println(isMatch)
}
