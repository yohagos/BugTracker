package models

import (
	"../utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Lastname  string             `bson:"lastname,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Password  string             `bson:"password,omitempty"`
	CreatedAt string             `bson:"createdAt,omitempty"`
	UpdatedAt string             `bson:"updatedAt,omitempty"`
}

/* // CreateTestUser func
func CreateTestUser() {
	var user bson.D
	time := utils.CreateTimeStamp()

	user = bson.D{
		{Key: "name", Value: "Yosef"},
		{Key: "lastname", Value: "Hagos"},
		{Key: "email", Value: "yosef@test.de"},
		{Key: "password", Value: "12345"},
		{Key: "createdAt", Value: time},
		{Key: "updatedAt", Value: time},
	}
	err := databases.AddNewUser(user)
	if err != nil {
		log.Fatal(err)
	}
} */

// CreateNewUser func
func CreateNewUser(newUser User) User {
	var user bson.D
	time := utils.CreateTimeStamp()

	user = bson.D{
		{Key: "name", Value: newUser.Name},
		{Key: "lastname", Value: newUser.Lastname},
		{Key: "email", Value: newUser.Email},
		{Key: "password", Value: newUser.Password},
		{Key: "createdAt", Value: time},
		{Key: "updatedAt", Value: time},
	}
	return user
}

func (user *User) GetUserID() string {
	return user.ID.String()
}

func (user *User) GetUserName() string {
	return user.Name
}

func (user *User) GetUserLastname() string {
	return user.Lastname
}

func (user *User) GetUserEmail() string {
	return user.Email
}

func (user *User) GetUserPassword() string {
	return user.Password
}

func (user *User) GetUserCreatedAt() string {
	return user.CreatedAt
}

func (user *User) GetUserUpdatedAt() string {
	return user.UpdatedAt
}
