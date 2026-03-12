package main

import (
	"fmt"
	"os"
)

func readBalanceFromFile() float64 {
	file, err := os.Open("balance.txt")
	if err != nil {
		if os.IsNotExist(err) {
			return 0
		}
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()
	var balance float64
	_, err = fmt.Fscanf(file, "%f", &balance)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return 0
	}
	return balance
}

func writeBalanceToFile(balance float64) {
	file, err := os.Create("balance.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("%.2f", balance))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func main() {
	var accountBalance float64 = readBalanceFromFile()

	for {
		presentOptions()
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your account balance is $%.2f\n", accountBalance)
		case 2:
			fmt.Println("Enter the amount to deposit: ")
			var amount float64
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Invalid amount. Please enter a positive number.")
				continue
			}
			accountBalance += amount
			fmt.Printf("Your new account balance is $%.2f\n", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Println("Enter the amount to withdraw: ")
			var amount float64
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Invalid amount. Please enter a positive number.")
				continue
			}
			if amount > accountBalance {
				fmt.Println("Insufficient funds")
				continue
			}
			accountBalance -= amount
			fmt.Printf("Your new account balance is $%.2f\n", accountBalance)
			writeBalanceToFile(accountBalance)
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid input. Please enter a number between 1 and 4.")
		}
	}

}
