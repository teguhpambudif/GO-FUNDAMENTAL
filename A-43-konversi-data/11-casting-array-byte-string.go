package main

import "fmt"

func main() {
	var text1 = "halo"
	var b = []byte(text1)

	fmt.Printf("%d %d %d %d\n", b[0], b[1], b[2], b[3])

	var text2 = "teguh"
	var b2 = []byte(text2)
	fmt.Printf("%d %d %d %d %d\n", b2[0], b2[1], b2[2], b2[3], b2[4])

	var byte1 = []byte{104, 97, 108, 111}
	var s = string(byte1)

	fmt.Printf("%s\n", s)
}
