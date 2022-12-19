package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	var text = "secret_text101"
	var sha = sha1.New()
	sha.Write([]byte(text))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)

	fmt.Println(encryptedString)

	check(encryptedString, "f4ebfd7a42d9a43a536e2bed9ee4974abf8f8dc8")
}

func check(a string, b string) {
	if a == b {
		fmt.Println("sama")
	} else {
		fmt.Println("tidaak")
	}
}
