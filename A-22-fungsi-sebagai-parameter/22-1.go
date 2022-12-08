package main

import (
	"fmt"
	"strings"
)

/*
ALIAS-ing on closure schema,
type FilterCallback func(string) bool

func filter(data []string, callback FilterCallback) []string {
    // ...
}

*/

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

func main() {
	var data = []string{"wick", "jason", "ethan", "kulo", "teguh", "tegoh"}
	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "o") // this is a function on "strings"
	})
	var dataLenght5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli \t\t:", data)
	// data asli
	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
	// filter ada huruf "o"
	fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)
	// filter jumlah huruf "5"
}
