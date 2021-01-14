package models

import (
	"context"
	"fmt"
	"log"

	"../apperrors"
	"../databases"
	"../utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ctx = context.TODO()
)

// BugTypes struct
type BugTypes struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Description string             `bson:"description,omitempty"`
	Acronym     string             `bson:"acronym,omitempty"`
	Name        string             `bson:"name,omitempty"`
	CreatedAt   string             `bson:"createdAt,omitempty"`
	UpdatedAt   string             `bson:"updatedAt,omitempty"`
}

// GetBugTypeID method
func (bugtype *BugTypes) GetBugTypeID() string {
	return bugtype.ID.String()
}

// GetBugTypeDescription method
func (bugtype *BugTypes) GetBugTypeDescription() string {
	return bugtype.Description
}

// GetBugTypeAcronym method
func (bugtype *BugTypes) GetBugTypeAcronym() string {
	return bugtype.Acronym
}

// GetBugTypeName method
func (bugtype *BugTypes) GetBugTypeName() string {
	return bugtype.Name
}

// GetBugTypeCreatedAt method
func (bugtype *BugTypes) GetBugTypeCreatedAt() string {
	return bugtype.CreatedAt
}

// GetBugTypeUpdatedAt method
func (bugtype *BugTypes) GetBugTypeUpdatedAt() string {
	return bugtype.UpdatedAt
}

func (bugtype *BugTypes) setBugTypeDescription(bt string) {
	bugtype.Description = bt
}

func (bugtype *BugTypes) setBugTypeAcronym(bt string) {
	bugtype.Acronym = bt
}

func (bugtype *BugTypes) setBugTypeName(bt string) {
	bugtype.Name = bt
}

func (bugtype *BugTypes) setBugTypeCreatedAt(bt string) {
	bugtype.CreatedAt = bt
}

func (bugtype *BugTypes) setBugTypeUpdatedAt(bt string) {
	bugtype.UpdatedAt = bt
}

// CreateNewBugType method
func (bugtype *BugTypes) CreateNewBugType() error {
	ok := BugTypeExists(bugtype.GetBugTypeAcronym())
	if ok {
		return apperrors.ErrorBugTypeAlreadyExists
	}

	timestamp := utils.CreateTimeStamp()
	bugDocument := bson.D{
		{Key: "description", Value: bugtype.Description},
		{Key: "acronym", Value: bugtype.Acronym},
		{Key: "name", Value: bugtype.Name},
		{Key: "createdAt", Value: timestamp},
		{Key: "updatedAt", Value: timestamp},
	}

	return databases.CreateNewBugType(bugDocument)

}

// BugTypeExists method
func BugTypeExists(acronym string) bool {
	return databases.CheckBugTypeExists(acronym)
}

// TestCreateNewBugType func
func TestCreateNewBugType() {
	time := utils.CreateTimeStamp()
	bugTypeDocument := bson.D{
		{Key: "description", Value: "Bugs were found which should be fixed before the next release."},
		{Key: "acronym", Value: "BUG"},
		{Key: "name", Value: "Bugs"},
		{Key: "createdAt", Value: time},
		{Key: "updatedAt", Value: time},
	}
	databases.CreateNewBugType(bugTypeDocument)
}

// NewBugTypeExists method
func NewBugTypeExists(acronym string) error {
	if err := databases.BugTypeCollection.FindOne(ctx, bson.M{"acronym": acronym}); err != nil {
		log.Println(err)
		return apperrors.ErrorBugTypeAlreadyExists
	}
	return nil
}

// BugTypeGetAllInformations func
func BugTypeGetAllInformations(acronym string) (BugTypes, error) {
	var bugtype BugTypes

	result, err := databases.GetAllBugTypeInformations(acronym)
	if err != nil {
		log.Println(err)
		return bugtype, err
	}

	for k, v := range result {
		switch k {
		case "acronym":
			key := fmt.Sprintf("%v", v)
			bugtype.setBugTypeAcronym(key)
		case "description":
			key := fmt.Sprintf("%v", v)
			bugtype.setBugTypeDescription(key)
		case "name":
			key := fmt.Sprintf("%v", v)
			bugtype.setBugTypeName(key)
		case "createdAt":
			key := fmt.Sprintf("%v", v)
			bugtype.setBugTypeCreatedAt(key)
		case "updatedAt":
			key := fmt.Sprintf("%v", v)
			bugtype.setBugTypeUpdatedAt(key)
		default:

		}
	}
	return bugtype, nil
}

// BugTypeListOfAcronyms func
func BugTypeListOfAcronyms() ([]string, error) {
	list, err := databases.GetListOfAllBugTypes()
	if err != nil {
		log.Println(err)
		return list, err
	}
	list = removeDuplicates(list)
	return list, nil
}

func removeDuplicates(list []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for v := range list {
		if encountered[list[v]] == true {

		} else {
			encountered[list[v]] = true
			result = append(result, list[v])
		}
	}
	return result
}
