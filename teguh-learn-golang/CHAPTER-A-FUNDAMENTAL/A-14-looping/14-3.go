package main

import "fmt"

func main() {
	// without condition after keywod "for", it will be infinite loop, so need the "break" keyword
	var i = 0

	for {
		fmt.Println("Angka", i)
		i++
		if i == 5 {
			break
		}
	}
}
