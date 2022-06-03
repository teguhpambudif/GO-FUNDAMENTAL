package main

import "fmt"

/*
Definition Closure is a function that can be stored in a variable. By applying these concepts, we can create functions within functions, or even create functions that return functions.
*/

func main() {
	// closure saved as variable
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)
	fmt.Printf("data : %v\nmin : %v\nmax : %v\n", numbers, min, max)
}
