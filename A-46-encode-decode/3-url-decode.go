package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var url = "https://twitter.com/teguhpambudif/status/1476304642162106373"

	var encoded = make([]byte, base64.URLEncoding.EncodedLen(len(url)))
	base64.URLEncoding.Encode(encoded, []byte(url))
	var encodedUrl = string(encoded)
	fmt.Println(encodedUrl)

	var decoded = make([]byte, base64.URLEncoding.DecodedLen(len(encoded)))
	var _, err = base64.URLEncoding.Decode(decoded, encoded)
	if err != nil {
		fmt.Println(err.Error())
	}
	var decodedString = string(decoded)
	fmt.Println(decodedString)

}
