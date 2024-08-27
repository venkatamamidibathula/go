package main

import (
	"fmt"
	"os"
	"strconv"
)

func getbalancefromfile() float64 {
	data, err := os.ReadFile("balance.txt")
	if err != nil {
		return 0
	}
	balance, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 0
	}
	return balance
}
func writedatatofile(balance float64) {
	balancedtext := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balancedtext), 0644)
}
func main() {
	var accountBalance float64
	for {
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Check Balance")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter the amount to deposit: ")
			var amount float64
			fmt.Scan(&amount)
			accountBalance += amount
			fmt.Println("Amount deposited successfully")
			writedatatofile(accountBalance)
		case 2:
			fmt.Print("Enter the amount to withdraw: ")
			var amount float64
			fmt.Scan(&amount)
			if amount > accountBalance {
				fmt.Println("Insufficient balance")
			} else {
				accountBalance -= amount
				fmt.Println("Amount withdrawn successfully")
			}
			writedatatofile(accountBalance)
		case 3:
			fmt.Println("Account Balance:", accountBalance)
		case 4:
			fmt.Println("Thank you for using GO Bank")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
