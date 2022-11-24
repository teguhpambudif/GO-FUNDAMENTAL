package main

import (
	"fmt"
	"strings"
)

func main() {
	var isExist = strings.Contains("John Wick", "Wick")
	fmt.Println(isExist)
}
