package main

import (
	"first-program/email/email"
	"fmt"
)

func main() {
	fmt.Println(email.IsValidEmail("test@test.com"))
	fmt.Println(email.IsValidEmail("test"))
	fmt.Println(email.IsValidEmail("test@"))
}
