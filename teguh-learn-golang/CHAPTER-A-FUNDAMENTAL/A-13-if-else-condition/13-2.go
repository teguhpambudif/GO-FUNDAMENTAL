package main

import "fmt"

func main() {
	var point = 8840.0

	// percent below is a temporary variable, only valid using ":=", and not "var"
	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s percent!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

}
