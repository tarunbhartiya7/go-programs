package main

import "fmt"

func getUserInput(userText string) float64 {
	var userInput float64

	fmt.Print(userText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateProfit(revenue float64, expenses float64, taxRate float64) (profit float64) {
	profit = revenue - expenses - taxRate
	return profit
}

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax rate: ")

	profit := calculateProfit(revenue, expenses, taxRate)
	fmt.Println("Profit: ", profit)
}
