package main

import "fmt"

func main() {
	orderSomeFood("pizza")
	orderSomeFood("burger")
}

func orderSomeFood(menu string) {
	defer fmt.Println("Terimakasih, silakan tunggu...")
	if menu == "pizza" {
		fmt.Print("Pilihan tepat!", " ")
		fmt.Print("Pizza di tempat kami paling enak!", "\n")
		return
	}

	fmt.Println("Pesanan anda:", menu)
}
