package main

import "fmt"

func main() {
	// map is a key-value pair, the syntax is using word map[key_type]value_type
	var chicken map[string]int
	chicken = map[string]int{} // bcuz default value map is nil, we need to initiate first using "{}" at the end of the type

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"])
	fmt.Println("februari", chicken["februari"])
	// fmt.Println("mei", chicken["mei"])	// bcuz mei isnt saved on map yet, then it will return default which is 0
}
