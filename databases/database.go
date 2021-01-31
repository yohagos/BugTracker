package databases

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx = context.TODO()

	// UserCollection exported MongoDB Collection
	UserCollection *mongo.Collection
	// BugTypeCollection exported MongoDB Collection
	BugTypeCollection *mongo.Collection
	// TicketCollection exported MongoDB Collection
	TicketCollection *mongo.Collection

	// VerificationCollection exported MongoDB Collection
	VerificationCollection *mongo.Collection

	mongoClient *mongo.Client
)

// Init func
func Init() {
	ClientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoClient, err := mongo.Connect(ctx, ClientOptions)
	if err != nil {
		log.Println(err)
	}

	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Connected to MongoDB!")

	UserCollection = mongoClient.Database("bugTracker").Collection("users")
	BugTypeCollection = mongoClient.Database("bugTracker").Collection("bugtype")
	TicketCollection = mongoClient.Database("bugTracker").Collection("tickets")
	VerificationCollection = mongoClient.Database("bugTracker").Collection("verification")

	/* listDatabases() */
}

/* func BsonToMapConvertor(document bson.M) map[string]string {
	var user map[string]string
	for _, v := range document {
		user = append(v, user)
	}
	return user
} */

func listDatabases() {
	databases, err := mongoClient.ListDatabases(ctx, bson.M{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(databases)
}

/*
func quickEntry() {
	_, err := quickCollection.InsertOne(ctx, bson.D{
		{Key: "title", Value: "Monkey King"},
		{Key: "author", Value: "Benjamin"},
	})
	if err != nil {
		log.Fatalln(err)
	}
} */

/* func TestBsonToMap() {
	var b bson.M
	if err := UserCollection.FindOne(ctx, bson.M{"email": "yosef.hagos@googlemail.com"}).Decode(&b); err != nil {
		log.Fatal(err)
	}
	//var m map[string]string
	for _, v := range b {
		log.Println(v)
	}
} */

/* // TestUser func
func TestUser() {
	time := utils.CreateTimeStamp()
	_, err := UserCollection.InsertOne(ctx, bson.D{
		{Key: "name", Value: "Yosef"},
		{Key: "lastname", Value: "Hagos"},
		{Key: "email", Value: "yosef.hagos@googlemail.com"},
		{Key: "password", Value: "12345"},
		{Key: "createdAt", Value: time},
		{Key: "updatedAt", Value: time},
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Created Test User !")
} */

/* // AddNewUser func
func AddNewUser(newUser bson.D) error {
	result, err := UserCollection.InsertOne(ctx, newUser)
	if err != nil {
		return err
	}
	//fmt.Printf("New User added: %v\n", result.InsertedID)
	return nil
} */

/* // UserAuthentification func
func UserAuthentification(username, password string) error {
	var user models.User
	if err := userCollection.FindOne(ctx, bson.M{"email": username}).Decode(&user); err != nil {
		log.Fatal(err)
	}
	if user.GetUserPassword() == password {
		return nil
	}
	ErrorUserDoesNotExist := errors.New("Login is invalid. Username / Password does not exists")

	return ErrorUserDoesNotExist
} */

/* // UserExists func
func UserExists(username string) bool {
	var user models.User
	if err := userCollection.FindOne(ctx, bson.M{"email": username}).Decode(&user); err != nil {
		log.Println(err)
		return false
	}
	return true
} */

/* func quickEntryTwo() {
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
	cursor, err := UserCollection.Find(ctx, bson.M{})
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
} */

/* // GetUserInformations func
func GetUserInformations(username string) *models.User {
	var document models.User
	if err := userCollection.FindOne(ctx, bson.M{"email": username}).Decode(&document); err != nil {
		log.Fatal(err)
	}
	return &document
} */

/* func readOneDocument() {
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

/* Replacing in MongoDB - es ist kein UPDATE, sondern ERSETZT das komplette Dokument
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

func dropCollection() {
	if err := UserCollection.Drop(ctx); err != nil {
		log.Fatal(err)
	}
} */
