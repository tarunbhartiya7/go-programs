package main

import (
	"fmt"
	"slices"
)

type Product struct {
	Name  string
	Price int
}

func Filter[T any](arr []T, filter func(T) bool) []T {
	var result []T
	for _, v := range arr {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	numbers := []int{5, 12, 8, 20, 3}
	var result []int
	for _, n := range numbers {
		if n > 10 {
			result = append(result, n)
		}
	}
	fmt.Println(result)

	products := []Product{
		{"Laptop", 1000},
		{"Mouse", 50},
		{"Keyboard", 150},
	}
	var filtered []Product
	for _, p := range products {
		if p.Price > 100 {
			filtered = append(filtered, p)
		}
	}
	fmt.Println(filtered)

	// Using Filter function
	result = Filter(numbers, func(n int) bool {
		return n > 15
	})
	fmt.Println(result)

	filtered = Filter(products, func(p Product) bool {
		return p.Price > 200
	})
	fmt.Println(filtered)

	numbersCopy := numbers[:]
	numbersCopy[1] = 100
	fmt.Println("numbersCopy:", numbersCopy)
	fmt.Println("numbers:", numbers)

	// use slices.clone to create a copy of the array
	numbersCopy = slices.Clone(numbers)
	numbersCopy[1] = 200
	fmt.Println("numbersCopy:", numbersCopy)
	fmt.Println("numbers:", numbers)

	// use make to create a copy of the array
	numbersCopy = make([]int, len(numbers))
	copy(numbersCopy, numbers)
	numbersCopy[1] = 300
	fmt.Println("numbersCopy:", numbersCopy)
	fmt.Println("numbers:", numbers)

	moreNums := []int{400, 500, 600}
	combinedArray := append(numbersCopy, moreNums...)
	fmt.Println("combinedArray:", combinedArray)
}
