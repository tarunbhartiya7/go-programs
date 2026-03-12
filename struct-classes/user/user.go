package user

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	email     string
	createdAt time.Time
}

// constructor function
func New(firstName string, lastName string, email string) *User {
	return &User{firstName: firstName, lastName: lastName, email: email, createdAt: time.Now()}
}

// class method
func (u User) DisplayName() {
	fmt.Println("Name: ", u.firstName, u.lastName)
}
