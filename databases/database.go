package databases

import (
	"context"
	"fmt"
	"log"

	"../models"
	"../utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()
var (
	quickCollection *mongo.Collection

	userCollection *mongo.Collection

	mongoClient *mongo.Client
)

/* // Episode struct
type Episode struct {
	ID          primitive.ObjectID `bson:"_id,omitempty`
	Podcast     primitive.ObjectID `bson:"_id,omitempty`
	Title       string             `bson:"title,omitempty`
	Description string             `bson:"description,omitempty`
	Duration    int32              `bson:"duration,omitempty`
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

	userCollection = mongoClient.Database("bugTracker").Collection("users")
	//quickCollection = mongoClient.Database("bugTracker").Collection("quick")

	//createPodcastEntry()

	//quickEntry()
	//quickEntryTwo()
	//quickEntriesMany()
	//deletingSingleDocument()
	//deletingManyDocument()
	//updating()
	//updatingMany()
	//ReadAllDocumentsFromQuickCollection()
	//readAllDocumentsFromQuickCollectionWithIteration()
	//replacingInDocument()
	//readOneDocument()
	//readOneDocumentFilter()
	//dropCollection()
	//listDatabases()
}

func listDatabases() {
	databases, err := mongoClient.ListDatabases(ctx, bson.M{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(databases)
}

func quickEntry() {
	_, err := quickCollection.InsertOne(ctx, bson.D{
		{Key: "title", Value: "Monkey King"},
		{Key: "author", Value: "Benjamin"},
	})
	if err != nil {
		log.Fatalln(err)
	}
}

// AddNewUser func
func AddNewUser(newUser bson.D) error {
	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return err
	}
	fmt.Printf("New User added: %v\n", result.InsertedID)
	return nil
}

func quickEntryTwo() {
	_, err := quickCollection.InsertOne(ctx, bson.D{
		{Key: "title", Value: "Monkey King"},
		{Key: "author", Value: "Benjamin"},
		{Key: "tags", Value: bson.A{"Monkey", "Sandy", "Pigsy"}},
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func quickEntriesMany() {
	_, err := quickCollection.InsertMany(ctx, []interface{}{
		bson.D{
			{Key: "title", Value: "Supernatural"},
			{Key: "author", Value: "Dean"},
			{Key: "description", Value: "Bro's kämpfen gegen Dämonen, Engel, mystische  Wesen und Gott"},
			{Key: "duration", Value: 250},
		},
		bson.D{
			{Key: "title", Value: "The Flash"},
			{Key: "author", Value: "Barry"},
			{Key: "description", Value: "The fastest Man alive..."},
			{Key: "duration", Value: 200},
		},
		bson.D{
			{Key: "title", Value: "Arrow"},
			{Key: "author", Value: "Oliver"},
			{Key: "description", Value: "Next Ras Al Ghul"},
			{Key: "duration", Value: 300},
		},
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func ReadAllDocumentsFromQuickCollection() {
	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalln(err)
	}

	var list []bson.M
	if err = cursor.All(ctx, &list); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(list)
}

func readAllDocumentsFromQuickCollectionWithIteration() {
	cursor, err := quickCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalln(err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var document bson.M
		if err = cursor.Decode(&document); err != nil {
			log.Fatalln(err)
		}
		fmt.Println(document)
	}
}

func GetUserInformations() {
	var document models.User
	if err := userCollection.FindOne(ctx, bson.M{"name": "Yosie"}).Decode(&document); err != nil {
		log.Fatal(err)
	}
	fmt.Println(document.Name)
}

func readOneDocument() {
	var document bson.M
	if err := quickCollection.FindOne(ctx, bson.M{}).Decode(&document); err != nil {
		log.Fatal(err)
	}
	fmt.Println(document)
}

func readOneDocumentFilter() {
	filterCursor, err := quickCollection.Find(ctx, bson.M{"author": "Barry"})
	if err != nil {
		log.Fatalln(err)
	}
	defer filterCursor.Close(ctx)
	var document []bson.M
	if err = filterCursor.All(ctx, &document); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(document)
}

func sortingDocuments() {
	opts := options.Find()
	opts.SetSort(bson.D{{Key: "duration", Value: -1}})

	sortCursor, err := quickCollection.Find(ctx, bson.D{{Key: "duration", Value: bson.D{{"$gt", 24}}}}, opts)
	if err != nil {
		log.Fatalln(err)
	}

	var document []bson.M
	if err = sortCursor.All(ctx, &document); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(document)
}

func updating() {
	id, _ := primitive.ObjectIDFromHex("5fb6816b6a15dd3065886ead")
	result, err := quickCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{Key: "author", Value: "Castiel"}}},
		},
	)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
}

func updatingMany() {
	result, err := quickCollection.UpdateOne(
		ctx,
		bson.M{"titlt": "Arrow"},
		bson.D{
			{"$set", bson.D{{Key: "author", Value: "Diggle"}}},
		},
	)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
}

/* Replacing in MongoDB - es ist kein UPDATE, sondern ERSETZT das komplette Dokument */
func replacingInDocument() {
	result, _ := quickCollection.ReplaceOne(
		ctx,
		bson.M{"author": "Castiel"},
		bson.M{
			"title":  "Lucifer",
			"author": "Gott",
		},
	)
	fmt.Printf("Replaced %v Documents", result.ModifiedCount)
}

func deletingSingleDocument() {
	result, err := quickCollection.DeleteOne(ctx, bson.M{"title": "Arrow"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DeleteOne removed %v documents\n", result.DeletedCount)
}

func deletingManyDocument() {
	result, err := quickCollection.DeleteMany(ctx, bson.M{"duration": 200})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DeleteMany removed %v document(s)\n", result.DeletedCount)
}

func DropCollection() {
	if err := userCollection.Drop(ctx); err != nil {
		log.Fatal(err)
	}
}
