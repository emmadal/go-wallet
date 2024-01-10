package operation

import (
	"fmt"
	"os"
)

// Withdraw money
func Withdraw (amount float64) (float64, error) {
	var total float64
	value, _ := GetBalance("balance.txt")
	if amount <= 0 {
		fmt.Println("Invalid amount. Should be greater than 0")
	}  else if amount > value {
		fmt.Println("Sorry, you don't have enough money in your account")
	} else {
		total = value - amount
		content := fmt.Sprint(total)
		os.WriteFile("balance.txt", []byte(content), 0644)
	}
	return total, nil
}