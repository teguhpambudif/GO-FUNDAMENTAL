package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "teguh pambudi"

	var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encoded :", encodedString)

	var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
	var DecodedString = string(decodedByte)
	fmt.Println(DecodedString)
}
