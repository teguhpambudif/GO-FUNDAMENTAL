package main

import (
	"fmt"
	"runtime"
)

func sendMessage(ch chan<- string) { // channel direction, send data
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}
	close(ch)
}

func printMessage(ch <-chan string) { // channel direction, receive data
	for message := range ch {
		fmt.Println(message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)
	go sendMessage(messages)
	printMessage(messages)
}
