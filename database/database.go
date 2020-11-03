package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"../models"
	"../utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()
var (
	userCollection *mongo.Collection
	mongoClient    *mongo.Client
)

/* // Person struct
type Person struct {
	name string `bson:"name" json:"name"`
	city string `bson:"city" json:"city"`
	/* 	createdAt time.Time `bson:"createdAt" json:"createdAt"`
	   	updatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
} */

// Init func
func Init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	utils.IsError(err)

	err = client.Ping(context.TODO(), nil)
	utils.IsError(err)
	mongoClient = client
	fmt.Println("Connected to MongoDB!")

	userCollection = mongoClient.Database("bugTracker").Collection("user")
}

// CreateUser func
func CreateUser(user *models.User) error {
	/* _, err := collection.InsertOne(ctx, person)
	return err */
	log.Println(user)
	_, err := userCollection.InsertOne(ctx, bson.D{
		{Key: "name", Value: user.GetUserName()},
		{Key: "lastname", Value: user.GetUserLastname()},
		{Key: "email", Value: user.GetUserEmail()},
		{Key: "password", Value: user.GetUserPassword()},
		{Key: "createdAt", Value: time.Now()},
		{Key: "updatedAt", Value: time.Now()},
	})
	return err
}

/* // CreatePerson func
func CreatePerson(name, lastname, email,  string) {
	p := Person{name: name, city: city}
	if err := CreateDocument(p); err != nil {
		log.Fatal(err)
	}
} */

/* collection = mongoClient.Database("mydb").Collection("person")
yosef := person{
	ID:        primitive.NewObjectID(),
	name:      "Yosef",
	city:      "Darmstadt",
	createdAt: time.Now(),
	updatedAt: time.Now(),
}

_, err := collection.InsertOne(ctx, bson.M{
	"name": yosef.name,
})
ifError(err) */
