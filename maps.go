// package main

// import "fmt"

// func main() {
// 	m := make(map[string]int)
// 	m["apple"] = 1
// 	m["banana"] = 2
// 	m["cherry"] = 3
// 	fmt.Println(m)

// 	delete(m, "apple")
// 	fmt.Println(m)

// 	for key, value := range m {
// 		fmt.Println(key, " - ", value)
// 	}

// 	a := m
// 	a["apple"] = 10
// 	fmt.Println(a)
// 	fmt.Println(m)

// 	// check if tomato is in the products
// 	tomato, isFound := m["Tomato"]
// 	if isFound {
// 		fmt.Println("Tomato found: ", tomato)
// 	} else {
// 		fmt.Println("Tomato not found")
// 	}
// }
