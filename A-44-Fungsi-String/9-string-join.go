package main

import (
	"fmt"
	"strings"
)

func main() {
	var data = []string{"banana", "papaya", "tomato"}
	var str = strings.Join(data, "-")
	fmt.Println(str)

	var strLower = strings.ToLower("alAy")
	fmt.Println(strLower)

	var strToUpper = strings.ToUpper("eat!")
	fmt.Println(strToUpper)
}
