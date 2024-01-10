package main

import (
	"fmt"
	"os"
)

// CreateWallet allow to create your .txt wallet
func CreateWallet(filename string, content float64) {
	fileContent := fmt.Sprint(content)
	_, err := os.Stat(filename)
	if err != nil {
		os.WriteFile(filename, []byte(fileContent), 0644)
		return
	}
}
