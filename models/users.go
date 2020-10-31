package models

import (
	"fmt"
)

// User struct
type User struct {
	name     string
	lastname string
	email    string
	password string
}

// GetAllUserInformation func
func (user *User) GetAllUserInformation() *User {
	return user
}

// CreateNewUser func
func CreateNewUser(userList ...string) *User {
	fmt.Println("\nCreateNewUser")
	fmt.Println("Name: " + userList[0])
	fmt.Println("Lastname: " + userList[1])
	fmt.Println("Email: " + userList[2])
	fmt.Println("Password: " + userList[3] + "\n")

	user := User{name: userList[0], lastname: userList[1], email: userList[2], password: userList[3]}
	return &user
}

// GetUserName func
func (user *User) GetUserName() string {
	return user.name
}

// GetUserLastname func
func (user *User) GetUserLastname() string {
	return user.lastname
}

// GetUserEmail func
func (user *User) GetUserEmail() string {
	return user.email
}

// GetUserPassword func
func (user *User) GetUserPassword() string {
	return user.password
}

func helloWorld() {
	fmt.Println("HelloWorld")
}
