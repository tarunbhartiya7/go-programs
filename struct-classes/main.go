package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"example.com/app/user"
)

type str string

func (text str) log() {
	// wait for 1 second
	time.Sleep(1 * time.Second)
	// wait for 300 microseconds
	// time.Sleep(500 * time.Microsecond)
	fmt.Println("Extending the string type with a method", text)
}

func api() {
	// make a request to https://jsonplaceholder.typicode.com/posts
	// and return the response
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts/5")
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func main() {
	user := user.New("John", "Doe", "john.doe@example.com")
	user.DisplayName()

	s := str("Hello, World!")
	s.log()

	api()
}
