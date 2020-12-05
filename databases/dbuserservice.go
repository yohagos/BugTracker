package databases

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// CheckUserExists func
func CheckUserExists(username string) bool {
	if err := UserCollection.FindOne(ctx, bson.M{"email": username}); err != nil {
		return false
	}
	return true
}

// CreateNewUser func
func CreateNewUser(user bson.D) {
	_, err := UserCollection.InsertOne(ctx, user)
	if err != nil {
		log.Println(err)
	}
}
