package main

import (
	"fmt"
	"time"
)

func main() {
	var timer = time.NewTimer(4 * time.Second)
	fmt.Println("start")
	<-timer.C
	fmt.Println("finish")
}
