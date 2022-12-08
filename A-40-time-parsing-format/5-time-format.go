package main

import (
	"fmt"
	"time"
)

func main() {
	var date, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")

	var dateS1 = date.Format("Monday 02, January 2006 15:04 MST")
	fmt.Println("dateS1", dateS1)
	fmt.Printf("%T\n", dateS1)

	var dateS2 = date.Format(time.RFC3339)
	fmt.Println("dateS2", dateS2)
	fmt.Printf("%T", dateS2)
}
