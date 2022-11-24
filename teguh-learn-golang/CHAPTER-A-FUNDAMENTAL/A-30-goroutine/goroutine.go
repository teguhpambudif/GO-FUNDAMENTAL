package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(2) // how many core activated for executed the program

	go print(5, "halo") // new goroutine need to add go at the beginning
	print(5, "apa kabar")

	var input string
	fmt.Scanln(&input)
}
