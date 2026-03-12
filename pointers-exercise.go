// package main

// import "fmt"

// const (
// 	Active   = true
// 	Inactive = false
// )

// type Product struct {
// 	Name string
// 	Tag  bool
// }

// func Activate(tag *bool) {
// 	*tag = Active
// }

// func Inactivate(tag *bool) {
// 	*tag = Inactive
// }

// func Act(product Product) Product {
// 	product.Tag = Active
// 	return product
// }

// func main() {
// 	products := []Product{
// 		{Name: "Laptop", Tag: Active},
// 		{Name: "Smartphone", Tag: Active},
// 		{Name: "Tablet", Tag: Active},
// 		{Name: "Monitor", Tag: Active},
// 	}

// 	Inactivate(&products[1].Tag)
// 	Inactivate(&products[2].Tag)

// 	fmt.Println(products)

// 	// not the right way but just for understanding how structs are pass by value in Go
// 	products[1] = Act(products[1])
// 	fmt.Println(products)

// }
