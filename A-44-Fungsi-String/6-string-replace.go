package main

import (
	"fmt"
	"strings"
)

func main() {
	// Fungsi ini digunakan untuk replace atau mengganti bagian dari string dengan string tertentu. Jumlah substring yang di-replace bisa ditentukan, apakah hanya 1 string pertama, 2 string, atau seluruhnya.
	var text = "banana"
	var find = "a"
	var replaceWith = "o"

	var newText1 = strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1)

	var newText2 = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2)

	var newText3 = strings.Replace(text, find, replaceWith, -1) //Angka -1 akan menjadikan proses replace berlaku pada semua substring.
	fmt.Println(newText3)
}
