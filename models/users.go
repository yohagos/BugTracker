package models

import (
	"time"
)

// User struct
type User struct {
	name      string    `bson:"name" json:"name"`
	lastname  string    `bson:"lastname" json:"lastname"`
	email     string    `bson:"email" json:"email"`
	password  string    `bson:"password" json:"password"`
	createdAt time.Time `bson:"createdAt" json:"createdAt"`
	updatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

// GetAllUserInformation func
func (user *User) GetAllUserInformation() *User {
	return user
}

// CreateNewUser func
func CreateNewUser(userList ...string) *User {
	user := User{name: userList[0], lastname: userList[1], email: userList[2], password: userList[3], createdAt: time.Now(), updatedAt: time.Now()}
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

// GetUserCreatedAt func
func (user *User) GetUserCreatedAt() string {
	toString := user.createdAt.Format("2020-01-01 13:00:02")
	return toString
}

// GetUserUpdatedAt func
func (user *User) GetUserUpdatedAt() string {
	toString := user.updatedAt.Format("2020-01-01 13:00:02")
	return toString
}
