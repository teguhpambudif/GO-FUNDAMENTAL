package main

import "fmt"

func main() {
	// array without capacity element
	var numbers = [...]int{2, 3, 2, 4, 3} // the capacity will automatically be 5

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah element \t:", len(numbers))

}
