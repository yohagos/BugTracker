package databases

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// CreateNewBugType func
func CreateNewBugType(bugtype bson.D) {
	_, err := BugTypeCollection.InsertOne(ctx, bugtype)
	if err != nil {
		log.Println(err)
	}
}

// CheckBugTypeExists func
func CheckBugTypeExists(acronym string) bool {
	if err := BugTypeCollection.FindOne(ctx, bson.M{"acronym": acronym}); err != nil {
		return false
	}
	return true
}

// GetAllBugTypeInformations func
func GetAllBugTypeInformations(acronym string) (bson.M, error) {
	var result bson.M
	if err := BugTypeCollection.FindOne(ctx, bson.M{"acronym": acronym}).Decode(&result); err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}
