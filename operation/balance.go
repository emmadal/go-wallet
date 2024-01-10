package operation

import (
	"os"
	"strconv"
)

// GetBalance get the balance from file
func GetBalance (filename string) (float64, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	value := string(data)
	balance, err := strconv.ParseFloat(value, 64)
	return balance, nil
}