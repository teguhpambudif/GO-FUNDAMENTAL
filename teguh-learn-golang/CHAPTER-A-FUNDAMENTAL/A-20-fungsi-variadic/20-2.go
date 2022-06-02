package main

import "fmt"

func main() {
	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	var avg = calculate(numbers...) // in this one, parameter for calculate function is an array. it can be written by adding "..." after the name of the variable that is used as a param
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)

	fmt.Println(msg)
}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	var avg = float64(total) / float64(len(numbers))
	return avg
}
