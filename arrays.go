// package main

// import "fmt"

// type Product struct {
// 	Name  string
// 	Price int
// }

// func PrintTotalPriceAndItems(products [3]Product) {
// 	total := 0
// 	fmt.Println("Items in basket:")
// 	for i := 0; i < len(products); i++ {
// 		fmt.Println(products[i].Name, "-", products[i].Price)
// 		total += products[i].Price
// 	}
// 	fmt.Println("Total price: ", total)
// }

// func main() {
// 	// var arr [3]int
// 	// fmt.Println(arr)

// 	// arr := [...]int{5, 7, 45, 23}
// 	// fmt.Println(arr)
// 	// fmt.Println(arr[0])
// 	// fmt.Println(arr[1])
// 	// fmt.Println(arr[2])
// 	// fmt.Println(arr[3])

// 	// arr[0] = 100
// 	// fmt.Println(arr)
// 	// fmt.Println("Length of arr: ", len(arr))

// 	// for i := 0; i < len(arr); i++ {
// 	// 	if i == 1 {
// 	// 		continue
// 	// 	}
// 	// 	if i == 2 {
// 	// 		break
// 	// 	}
// 	// 	fmt.Println(arr[i])
// 	// }

// 	// // This is index out of range error
// 	// for i := 0; i < 10; i++ {
// 	// 	fmt.Println(arr[i])
// 	// }

// 	basket := [3]Product{
// 		{Name: "Apple", Price: 10},
// 		{Name: "Banana", Price: 20},
// 		{Name: "Cherry", Price: 30},
// 	}

// 	PrintTotalPriceAndItems(basket)

// }
