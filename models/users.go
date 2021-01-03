package models

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	"../apperrors"
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

func (user *User) setUserEmail(str string) {
	user.Email = str
}

func (user *User) setUserName(str string) {
	user.Name = str
}

func (user *User) setUserLastname(str string) {
	user.Lastname = str
}

func (user *User) setUserPassword(str string) {
	user.Password = str
}

func (user *User) setUserCreatedAt(str string) {
	user.CreatedAt = str
}

func (user *User) setUserUpdatedAt(str string) {
	user.UpdatedAt = str
}

// CreateNewUser func
func (user *User) CreateNewUser() error {
	ok := databases.CheckUserExists(user.GetUserName())
	if !ok {
		return apperrors.ErrorUserAlreadyExists
	}

	pwd := user.GetUserPassword()

	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), cost)

	if err != nil {
		return err
	}

	time := utils.CreateTimeStamp()
	userDocument := bson.D{
		{Key: "name", Value: user.Name},
		{Key: "lastname", Value: user.Lastname},
		{Key: "email", Value: user.Email},
		{Key: "password", Value: hash},
		{Key: "createdAt", Value: time},
		{Key: "updatedAt", Value: time},
	}
	databases.CreateNewUser(userDocument)
	return nil
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
func UserGetAllInformations(username string) (User, error) {
	var user User

	result, err := databases.GetAllUserInformations(username)
	if err != nil {
		log.Println(err)
		return user, err
	}

	for k, v := range result {
		switch k {
		case "name":
			key := fmt.Sprintf("%v", v)
			user.setUserName(key)
		case "lastname":
			key := fmt.Sprintf("%v", v)
			user.setUserLastname(key)
		case "email":
			key := fmt.Sprintf("%v", v)
			user.setUserEmail(key)
		case "password":
			key := fmt.Sprintf("%v", v)
			user.setUserPassword(key)
		case "createdAt":
			key := fmt.Sprintf("%v", v)
			user.setUserCreatedAt(key)
		case "updatedAt":
			key := fmt.Sprintf("%v", v)
			user.setUserUpdatedAt(key)
		default:

		}
	}

	return user, nil
}

// TestCreateUser func
func TestCreateUser() {
	timestamp := utils.CreateTimeStamp()
	userDocument := bson.D{
		{Key: "name", Value: "Yosef"},
		{Key: "lastname", Value: "Hagos"},
		{Key: "email", Value: "test@test.com"},
		{Key: "password", Value: "123456"},
		{Key: "createdAt", Value: timestamp},
		{Key: "updatedAt", Value: timestamp},
	}
	databases.CreateNewUser(userDocument)
}

func bsonToUser(list bson.M) User {
	var user User

	for k, v := range list {
		switch k {
		case "name":
			key := fmt.Sprintf("%v", v)
			user.setUserName(key)
		case "lastname":
			key := fmt.Sprintf("%v", v)
			user.setUserLastname(key)
		case "email":
			key := fmt.Sprintf("%v", v)
			user.setUserEmail(key)
		case "password":
			key := fmt.Sprintf("%v", v)
			user.setUserPassword(key)
		case "createdAt":
			key := fmt.Sprintf("%v", v)
			user.setUserCreatedAt(key)
		case "updatedAt":
			key := fmt.Sprintf("%v", v)
			user.setUserUpdatedAt(key)
		default:

		}
	}

	return user
}
