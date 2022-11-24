package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Digunakan untuk mencari index string kembalian hasil dari operasi regexp.
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	var idx = regex.FindStringIndex(text)
	fmt.Println(idx)

	var str = text[0:6]
	fmt.Println(str)
}
