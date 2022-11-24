package main

import "fmt"

func main() {
	// "break" will force stop loop, "continue" will force to the next loop or in other word skipping current true condition
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}
		if i > 8 {
			break
		}
		fmt.Println("Angka", i)
	}
}
