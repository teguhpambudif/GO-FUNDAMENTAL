package main

import "fmt"

func main() {
	var fruits = []string{"apple", "grape", "banana", "melon"}

	var aFruits = fruits[0:3] // make new slice from fruits
	var bFruits = fruits[1:4] // make new slice from fruits

	var aaFruits = aFruits[1:2] // new slice from aFruits
	var bbFruits = bFruits[0:1] // mew slice from bFruits

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(bbFruits)
	fmt.Println("")

	bbFruits[0] = "pinnaple" // in this we change "grape" into "pinnaple"

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(bbFruits)

}
