package main

import (
	"fmt"
	"time"
)

func count(n int) {
	for i := 0; i < n; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}
}

func slowGreet(name string, done chan bool) {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello, ", name)
	done <- true
}

func main() {
	// go count(5)
	// go count(3)
	// fmt.Println("wait for the goroutine to finish")
	// time.Sleep(2000 * time.Millisecond)
	// fmt.Println("done")

	fmt.Println("Loading...")
	done := make(chan bool)
	go slowGreet("John", done)
	<-done
	fmt.Println("done")
}
