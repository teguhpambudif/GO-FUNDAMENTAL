package main

import (
	"fmt"
	"strings"
)

func main() {
	// Digunakan untuk mencari posisi indeks sebuah string (parameter kedua) dalam string (parameter pertama).
	var index1 = strings.Index("ethan hunt", "ha")
	fmt.Println(index1)

	var index2 = strings.Index("ethan hunt", "n")
	fmt.Println(index2)
}
