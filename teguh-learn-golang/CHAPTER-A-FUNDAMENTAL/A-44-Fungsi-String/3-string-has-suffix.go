package main

import (
	"fmt"
	"strings"
)

func main() {
	// Digunakan untuk deteksi apakah sebuah string (parameter pertama) diakhiri string tertentu (parameter kedua).
	var isSuffix1 = strings.HasSuffix("john wick", "ic")
	fmt.Println(isSuffix1)

	var isSuffix2 = strings.HasSuffix("john wick", "ck")
	fmt.Println(isSuffix2)
}
