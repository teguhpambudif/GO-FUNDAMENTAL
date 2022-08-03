package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("halo") // will not be executed
	os.Exit(1)
	fmt.Println("selamat datang")
}
