package models

import (
	"context"
	"log"

	"../apperrors"
	"../databases"

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

func test1() {
	log.Println(apperrors.ErrorUserDoesNotExist)
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

// CreateNewBugType method
func (bugtype *BugTypes) CreateNewBugType() error {
	ok := bugtype.BugTypeExists()
	if ok {
		log.Println(apperrors.ErrorBugTypeAlreadyExists)
		return apperrors.ErrorBugTypeAlreadyExists
	}
	return nil
}

// BugTypeExists method
func (bugtype *BugTypes) BugTypeExists() bool {
	if err := databases.BugTypeCollection.FindOne(ctx, bson.M{"acronym": bugtype.GetBugTypeAcronym()}); err != nil {
		log.Fatalln(err)
		return false
	}
	return true
}
