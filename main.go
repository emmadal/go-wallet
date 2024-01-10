package main

import (
	"fmt"
	"github.com/emmadal/gowallet/wallet/operation"
	"log"
)

// BalanceFile file name
const BalanceFile = "balance.txt"

func main() {

	fmt.Println("Welcome to GoWallet Bank!")
	CreateWallet(BalanceFile, 0)

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var amount float64 = 0
		var choice int

		fmt.Print("Enter your choice: ")
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			log.Fatalln(err)
		}

		switch choice {
		case 1:
			balance, _ := operation.GetBalance(BalanceFile)
			fmt.Println("🟢 Your account balance is", balance)
		case 2:
			fmt.Print("How much do you want to deposit? ")
			fmt.Scan(&amount)
			balance, _ := operation.Deposit(amount)
			fmt.Println("🎉 You new balance is ", balance)

		case 3:
			fmt.Print("How much do you want to withdraw? ")
			fmt.Scan(&amount)
			balance, _ := operation.Withdraw(amount)
			fmt.Println("🍺 Your new balance is", balance)
		case 4:
			fmt.Println("👋 Thanks for using the GoWallet CLI!")
			return
		default:
			fmt.Printf("%d is not a valid option.\n", choice)
		}
	}
}
