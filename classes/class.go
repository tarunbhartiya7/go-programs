package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Receiver function
// we are attaching the function to the Person struct like class methods in other languages
func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name, "and I am", p.Age, "years old")
}

func main() {
	p := Person{Name: "John", Age: 30}
	p.Greet()
}
