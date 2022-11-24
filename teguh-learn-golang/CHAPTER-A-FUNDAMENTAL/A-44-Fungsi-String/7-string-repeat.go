package main

import (
	"fmt"
	"strings"
)

func main() {
	// Digunakan untuk mengulang string (parameter pertama) sebanyak data yang ditentukan (parameter kedua).
	var str = strings.Repeat("ma", 4)
	fmt.Println(str)
}
