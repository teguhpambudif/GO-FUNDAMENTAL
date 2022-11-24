package main

import (
	"fmt"
	"strings"
)

func main() {
	// Digunakan untuk deteksi apakah sebuah string (parameter pertama) diawali string tertentu (parameter kedua).
	var isPrefix = strings.HasPrefix("john wick", "jo")
	fmt.Println(isPrefix)

	var isPrefix2 = strings.HasPrefix("john wick", "wi")
	fmt.Println(isPrefix2)
}
