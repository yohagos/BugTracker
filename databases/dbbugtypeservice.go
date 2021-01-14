package databases

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// CreateNewBugType func
func CreateNewBugType(bugtype bson.D) error {
	_, err := BugTypeCollection.InsertOne(ctx, bugtype)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
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

// GetListOfAllBugTypes func
func GetListOfAllBugTypes() ([]string, error) {
	cursor, err := BugTypeCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var btList []bson.M
	if err = cursor.All(ctx, &btList); err != nil {
		log.Println(err)
		return nil, err
	}
	return filterBugTypeList(btList), nil
}

func filterBugTypeList(bugtypeList []bson.M) []string {
	var list []string
	for _, ListValue := range bugtypeList {
		for key, unit := range ListValue {
			if key == "acronym" {
				val := fmt.Sprintf("%v", unit)
				list = append(list, val)
			}
		}
	}
	return list
}
