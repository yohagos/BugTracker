package models

import (
	"log"
	"time"

	"../databases"

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
	CreatedAt time.Time          `bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty"`
}

// CreateTestUser func
func CreateTestUser() {
	var user bson.D
	time := CreateTimeStamp()

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
}

func CreateTimeStamp() string {
	current_time := time.Now()
	return current_time.Format("2006-01-02 15:04:05")
}
