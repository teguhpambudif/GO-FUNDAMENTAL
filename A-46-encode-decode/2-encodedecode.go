package main

import (
	"encoding/base64"
	"fmt"
)

// hampir sama dengan toString, tapi variabel byte harus ditampung terlebih dulu dan panjang byte harus sesuai dengan hasil yang akan ditampung
func main() {
	var data = "teguh pambudi lagi"

	var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, []byte(data))
	var encodedString = string(encoded)
	fmt.Println(encodedString)

	var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	var _, err = base64.StdEncoding.Decode(decoded, encoded)
	if err != nil {
		fmt.Println(err.Error())
	}
	var decodedString = string(decoded)
	fmt.Println(decodedString)
}
