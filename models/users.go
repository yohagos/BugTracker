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

// GetUserID func
func (user *User) GetUserID() string {
	return user.ID.String()
}

// GetUserName func
func (user *User) GetUserName() string {
	return user.Name
}

// GetUserLastname func
func (user *User) GetUserLastname() string {
	return user.Lastname
}

// GetUserEmail func
func (user *User) GetUserEmail() string {
	return user.Email
}

// GetUserPassword func
func (user *User) GetUserPassword() string {
	return user.Password
}

// GetUserCreatedAt func
func (user *User) GetUserCreatedAt() string {
	return user.CreatedAt
}

// GetUserUpdatedAt func
func (user *User) GetUserUpdatedAt() string {
	return user.UpdatedAt
}

// CreateNewUser func
func (user *User) CreateNewUser() {
	ok := user.UserExists()
	if !ok {
		log.Println("Username already exists")
		return
	}

	time := utils.CreateTimeStamp()
	userDocument := bson.D{
		{Key: "name", Value: user.Name},
		{Key: "lastname", Value: user.Lastname},
		{Key: "email", Value: user.Email},
		{Key: "password", Value: user.Password},
		{Key: "createdAt", Value: time},
		{Key: "updatedAt", Value: time},
	}

	_, err := databases.UserCollection.InsertOne(ctx, userDocument)
	if err != nil {
		log.Fatalln(err)
	}
	//databases.AddNewUser(userDocument)
}

// UserExists func
func (user *User) UserExists() bool {
	if err := databases.UserCollection.FindOne(ctx, bson.M{"email": user.GetUserName()}).Decode(&user); err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// UserAuthentification func
func UserAuthentification(username, password string) error {
	var user User
	if err := databases.UserCollection.FindOne(ctx, bson.M{"email": username}).Decode(&user); err != nil {
		log.Println(err)
		return err
	}

	if user.GetUserPassword() == password {
		return nil
	}
	return utils.ErrorUserDoesNotExist
}
