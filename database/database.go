package database

import (
	"context"
	"fmt"
	"time"

	utils "../utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()
var (
	collection  *mongo.Collection
	mongoClient *mongo.Client
)

type person struct {
	ID        primitive.ObjectID `bson:"_id"`
	name      string             `bson:"name"`
	city      string             `bson:"city"`
	createdAt time.Time          `bson:"createdAt"`
	updatedAt time.Time          `bson:"updatedAt"`
}

func DatabaseInit() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	utils.IsError(err)

	err = client.Ping(context.TODO(), nil)
	utils.IsError(err)
	mongoClient = client
	fmt.Println("Connected to MongoDB!")
}

func createDocument(person person) error {
	_, err := collection.InsertOne(ctx, person)
	return err
}

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
