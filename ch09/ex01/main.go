package main

import (
	"fmt"

	"github.com/mikan/go-training-course/ch09/ex01/bank"
)

func main() {
	fmt.Printf("balance: %d\n", bank.Balance())

	// deposit
	fmt.Printf("deposit (n>0): ")
	var amount int
	if _, err := fmt.Scan(&amount); err != nil || amount <= 0 {
		return
	}
	bank.Deposit(amount)
	fmt.Printf("balance: %d\n", bank.Balance())

	for {
		// withdraw
		fmt.Printf("withdraw (n>0): ")
		if _, err := fmt.Scan(&amount); err != nil || amount <= 0 {
			return
		}
		if bank.Withdraw(amount) {
			fmt.Println("WITHDRAW SUCCESS")
			fmt.Printf("balance: %d\n", bank.Balance())
		} else {
			fmt.Println("WITHDRAW FAILED")
		}
	}
}
