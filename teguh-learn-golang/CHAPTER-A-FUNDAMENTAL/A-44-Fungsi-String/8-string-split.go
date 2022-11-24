package main

import (
	"fmt"
	"strings"
)

func main() {
	// Digunakan untuk memisah string (parameter pertama) dengan tanda pemisah bisa ditentukan sendiri (parameter kedua). Hasilnya berupa slice string.
	var string1 = strings.Split("the dark knight", " ")
	fmt.Println(string1)
	// for _, i := range string1 {
	// 	fmt.Println(i)
	// }

	var string2 = strings.Split("batman", "")
	fmt.Println(string2)
}
