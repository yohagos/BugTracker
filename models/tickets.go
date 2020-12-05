package models

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Tickets struct
type Tickets struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Lastname  string             `bson:"lastname,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Password  string             `bson:"password,omitempty"`
	CreatedAt string             `bson:"createdAt,omitempty"`
	UpdatedAt string             `bson:"updatedAt,omitempty"`
}

func test() {
	log.Println()
}
