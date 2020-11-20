package models

import (
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
func CreateNewUser(newUser User) {
	/* var user bson.D
	time := utils.CreateTimeStamp()

	user = bson.D{
		{Key: "name", Value: newUser.Name},
		{Key: "lastname", Value: newUser.Lastname},
		{Key: "email", Value: newUser.Email},
		{Key: "password", Value: newUser.Password},
		{Key: "createdAt", Value: time},
		{Key: "updatedAt", Value: time},
	}
	err := databases.AddNewUser(user)
	utils.IsError(err) */
}

/* func GetUser(){
	//objid := "5fb7e71afac4670b98688ae8"


} */
