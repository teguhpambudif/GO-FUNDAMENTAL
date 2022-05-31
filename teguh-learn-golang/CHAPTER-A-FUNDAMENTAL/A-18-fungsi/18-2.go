package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
}

// func below is to generate random number based on param we inout in main
func randomWithRange(min, max int) int { // "int" after (param, param int) is the return data type of the function
	var value = rand.Int()%(max-min+1) + min
	return value
}
