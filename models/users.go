package models

import (
	"log"

	"../databases"
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
	ok := databases.CheckUserExists(user.GetUserName())
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
	databases.CreateNewUser(userDocument)
}

// UserExists func
func UserExists(username string) bool {
	return databases.CheckUserExists(username)
}

// UserAuthentification func
func UserAuthentification(username, password string) error {
	err := databases.AuthentificationUser(username, password)
	return err
}

// UserGetAllInformations func
func UserGetAllInformations(username string) (*User, error) {
	//var user *User
	result, err := databases.GetAllUserInformations(username)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	user, _ := bson.Marshal(result)
	log.Println(user)
	return nil, nil
}

// TestCreateUser func
func TestCreateUser() {
	timestamp := utils.CreateTimeStamp()
	userDocument := bson.D{
		{Key: "name", Value: "user.Name"},
		{Key: "lastname", Value: "user.Lastname"},
		{Key: "email", Value: "user.Email"},
		{Key: "password", Value: "user.Password"},
		{Key: "createdAt", Value: timestamp},
		{Key: "updatedAt", Value: timestamp},
	}
	databases.CreateNewUser(userDocument)
}
