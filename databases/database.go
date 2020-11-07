package databases

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
	postCollection *mongo.Collection
	mongoClient    *mongo.Client
)

type Post struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
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

	userCollection = mongoClient.Database("bugTracker").Collection("user")
	postCollection = mongoClient.Database("bugTracker").Collection("post")

	/* title := "Whats up Buddy?"

	body := "How are you? If you like, we could visit Tokyo next week. Best Regards, Yosie"
	*/
	/* postToBen := Post{"Ben", "Whats Up"}
	postToMary := Post{"Mary", "Looking Good"}
	postToDean := Post{"Sam", "Family Business"}

	insertMany := []interface{}{postToBen, postToMary, postToDean} */
	/* id := insertPost(title, body) */
	/* insertManyPost(insertMany)

	findDocument("Sam") */
	//getPost(id)
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
func GetAllUsers() /* ([]bson.M, error) */ {
	/* var resultsList []bson.M */

	cursor, err := userCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
		defer cursor.Close(ctx)
		/* return nil, err */
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {

		fmt.Println(cursor.Current)

		var userResult models.User
		err := cursor.Decode(&userResult)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(userResult, "\n")
		/* var result interface{}
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		//fmt.Println(result)
		findUser := models.User{}
		b, err := json.Unmarshal(findUser, &result)
		fmt.Println(string(b))
		results = append(results, result) */
	}

	/* return results, err */
}

func insertOnePost(title, body string) string {
	post := Post{title, body}

	insertResults, err := postCollection.InsertOne(ctx, post)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted ID: ", insertResults.InsertedID)

	return fmt.Sprint(insertResults.InsertedID)
}

func insertManyPost(manyPosts []interface{}) {
	/* /_, err := postCollection.InsertMany(ctx, manyPosts)
	if err != nil {
		log.Fatal(err)
	} */

}

/* func getPost(id string) {
	filter := bson.D{{}}

	var post Post

	err := postCollection.FindOne(ctx, filter).Decode(&post)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found post with Title: ", post.Title)
} */

func FindDocument(who string) {
	filter := bson.D{{Key: "name", Value: who}}
	var result bson.D
	err := userCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range result {
		fmt.Println(v)
		for _, w := range v {
			fmt.Println(w)
		}
	}

}

func UpdateUser(name string) {
	filter := bson.D{{"name", "Yosie"}}
	newName := bson.D{
		{"$set", bson.D{
			{"name", name},
			{"updatedAt", time.Now()},
		}}}

	res, err := userCollection.UpdateOne(ctx, filter, newName)
	if err != nil {
		log.Fatal(err)
	}
	updatedObject := *res

	fmt.Printf("Matched count: %d\nModified count: %d\n", updatedObject.MatchedCount, updatedObject.ModifiedCount)
}
