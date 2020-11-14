package databases

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"../utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()
var (
	/* userCollection    *mongo.Collection
	postCollection    *mongo.Collection */
	podcastCollection *mongo.Collection
	episodeCollection *mongo.Collection

	quickCollection *mongo.Collection

	mongoClient *mongo.Client
)

// Podcast struct
type Podcast struct {
	ID     primitive.ObjectID `bson:"_id,omitempty`
	Title  string             `bson:"title,omitempty`
	Author string             `bson:"author,omitempty`
	Tags   []string           `bson:"tags,omitempty`
}

// Episode struct
type Episode struct {
	ID          primitive.ObjectID `bson:"_id,omitempty`
	Podcast     primitive.ObjectID `bson:"_id,omitempty`
	Title       string             `bson:"title,omitempty`
	Description string             `bson:"description,omitempty`
	Duration    int32              `bson:"duration,omitempty`
}

// Init func
func Init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	utils.IsError(err)

	err = client.Ping(context.TODO(), nil)
	utils.IsError(err)
	mongoClient = client
	fmt.Println("Connected to MongoDB!")

	/* userCollection = mongoClient.Database("bugTracker").Collection("user")
	postCollection = mongoClient.Database("bugTracker").Collection("post") */

	podcastCollection = mongoClient.Database("bugTracker").Collection("podcast")
	episodeCollection = mongoClient.Database("bugTracker").Collection("episode")

	quickCollection = mongoClient.Database("bugTracker").Collection("quick")

	//createPodcastEntry()

	//quickEntry()
	//quickEntryTwo()
	//quickEntriesMany()

	//readAllDocumentsFromQuickCollection()
	//readAllDocumentsFromQuickCollectionWithIteration()

	//readOneDocument()
	readOneDocumentFilter()

	//listDatabases()
}

func listDatabases() {
	databases, err := mongoClient.ListDatabases(ctx, bson.M{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(databases)
}

/*
func createPodcastEntry() {
	podcast := Podcast{
		Title:  "How to ? Golang & Mongo",
		Author: "Yosef Hagos",
		Tags:   []string{"development", "programming", "coding", "go", "golang", "mongodb"},
	}

	insertResult, err := podcastCollection.InsertOne(ctx, podcast)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(insertResult.InsertedID)
} */

func quickEntry() {
	_, err := quickCollection.InsertOne(ctx, bson.D{
		{Key: "title", Value: "Monkey King"},
		{Key: "author", Value: "Benjamin"},
	})
	if err != nil {
		log.Fatalln(err)
	}
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

func readAllDocumentsFromQuickCollection() {
	cursor, err := quickCollection.Find(ctx, bson.M{})
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
