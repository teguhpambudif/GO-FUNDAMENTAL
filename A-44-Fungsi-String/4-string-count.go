package main

import (
	"fmt"
	"strings"
)

func main() {
	// Memiliki kegunaan untuk menghitung jumlah karakter tertentu (parameter kedua) dari sebuah string (parameter pertama). Nilai kembalian fungsi ini adalah jumlah karakternya.
	var howMany = strings.Count("ethan hunt", "t")
	fmt.Println(howMany)
	fmt.Printf("%T\n", howMany)
}
