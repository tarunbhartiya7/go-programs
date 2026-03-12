// package main

// import "fmt"

// func PrintArrayElements(arr []int) {
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Println(arr[i])
// 	}
// }

// func main() {
// 	slice := []int{1, 2, 3, 4, 5}
// 	slice2 := slice[1:3]

// 	fmt.Println(slice)
// 	fmt.Println(slice2)

// 	slice = append(slice, 6)
// 	fmt.Println(slice)

// 	combined := append(slice, slice2...)
// 	fmt.Println(combined)

// 	assemblyLine := make([]string, 3)
// 	assemblyLine[0] = "Car"
// 	assemblyLine[1] = "Plane"
// 	assemblyLine[2] = "Train"
// 	fmt.Println(assemblyLine)
// }
