package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(10)
	fmt.Println("random ke-1:", rand.Int())
	fmt.Println("ramdom ke-2:", rand.Int())
	fmt.Println("random ke-3:", rand.Int())

}
