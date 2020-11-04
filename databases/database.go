package databases

import (
	"context"
	"fmt"
	"log"

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
	_, err := userCollection.InsertOne(ctx, bson.D{
		{Key: "name", Value: user.GetUserName()},
		{Key: "lastname", Value: user.GetUserLastname()},
		{Key: "email", Value: user.GetUserEmail()},
		{Key: "password", Value: user.GetUserPassword()},
		{Key: "createdAt", Value: user.GetUserCreatedAt()},
		{Key: "updatedAt", Value: user.GetUserUpdatedAt()},
	})
	return err
}

// GetAllUsers func
func GetAllUsers() ([]*models.User, error) {
	findOptions := options.Find()

	var results []*models.User

	cursor, err := userCollection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var element *models.User
		err := cursor.Decode(&element)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, element)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return results, err
}
