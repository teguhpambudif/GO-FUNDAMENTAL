package main

import "fmt"

func main() {
	dividedNumber(10, 2)
	dividedNumber(4, 0)
	dividedNumber(8, -4)
}

func dividedNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
		return // "return" will stop block function when condition true
	}
	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}
