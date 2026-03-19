package main

import "fmt"

// generator function using channel
func fibonacci(n int, ch chan int) {
	defer close(ch)

	a, b := 1, 2

	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
}

func main() {
	n := 10
	ch := make(chan int)

	// start goroutine
	go fibonacci(n, ch)

	// read from channel
	for num := range ch {
		fmt.Println(num)
	}
}
