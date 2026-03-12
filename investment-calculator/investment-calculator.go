package main

import (
	"fmt"
	"math"
)

func main() {
	var amount float64
	var rate float64
	var period int
	var inflationRate float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&amount)

	fmt.Print("Enter rate of interest (in percent): ")
	fmt.Scan(&rate)

	fmt.Print("Enter period (in years): ")
	fmt.Scan(&period)

	fmt.Print("Enter inflation rate (in percent): ")
	fmt.Scan(&inflationRate)

	// Using compound interest formula: A = P * (1 + r/n)^(nt)
	// For annual compounding, n = 1. Interest = A - P
	finalAmount := amount * math.Pow(1+rate/100, float64(period))
	inflationAdjustedAmount := amount * math.Pow(1+(rate-inflationRate)/100, float64(period))

	fmt.Printf("Total amount with interest: %.2f\n", math.Round(finalAmount))
	fmt.Printf("Total amount with interest and inflation: %.2f\n", math.Round(inflationAdjustedAmount))
}
