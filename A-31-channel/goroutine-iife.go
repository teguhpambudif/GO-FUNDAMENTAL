package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	var messages = make(chan string)

	go func(who string) { // new IIFE goroutine
		var data = fmt.Sprintf("hello %s", who)
		messages <- data
	}("wick")

	var mesage = <-messages
	fmt.Println(mesage)
}
