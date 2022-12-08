package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string) // messages is a channel

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data // var data stored to channel
	}

	go sayHelloTo("john wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("jason bourne")

	var message1 = <-messages // channel passing data to initialized variable
	fmt.Println(message1)

	var message2 = <-messages // channel passing data to initialized variable
	fmt.Println(message2)

	var message3 = <-messages // channel passing data to initialized variable
	fmt.Println(message3)
}
