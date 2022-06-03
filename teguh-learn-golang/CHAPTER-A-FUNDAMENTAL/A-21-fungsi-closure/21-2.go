package main

import "fmt"

// Immediately-Invoked Function Expression (IIFE), This type of closure is executed immediately upon declaration. Usually used to wrap processes that are only performed once, may return a value, or may not.

func main() {
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3) // ciri khas IIFE dengan adanya kurung parameter setekah closure berakhir
	fmt.Println("original number :", numbers)
	fmt.Println("filtered number:", newNumbers)
}
