package models

import (
	"context"
	"log"

	"../databases"
	"../utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ctx = context.TODO()

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

// CreateNewUser func
func (user *User) CreateNewUser() {
	time := utils.CreateTimeStamp()

	userDocument := bson.D{
		{Key: "name", Value: user.Name},
		{Key: "lastname", Value: user.Lastname},
		{Key: "email", Value: user.Email},
		{Key: "password", Value: user.Password},
		{Key: "createdAt", Value: time},
		{Key: "updatedAt", Value: time},
	}

	databases.AddNewUser(userDocument)
}

func (user *User) UserExists() bool {
	if err := databases.UserCollection.FindOne(ctx, bson.M{"email": user.GetUserName()}).Decode(&user); err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func UserAuthentification(username, password string) {
	/* err := databases.UserAuthentification(username, password)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil */

	// UserAuth muss noch angepasst werden
}
