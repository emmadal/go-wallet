package operation

import (
	"fmt"
	"os"
)

// Deposit function increase user balance
func Deposit (amount float64) (float64, error) {
	var total float64
	if amount <= 0 {
		fmt.Println("Invalid amount. Should be greater than 0")
	} else {
		wallet, err := GetBalance("balance.txt")
		if err != nil {
			return 0, err
		}
		total = wallet + amount
		content := fmt.Sprint(total)
		os.WriteFile("balance.txt", []byte(content), 0644)
	}
	return total, nil
}