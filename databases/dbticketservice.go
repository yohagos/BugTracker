package databases

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// CreateNewTicket func
func CreateNewTicket(ticket bson.D) {
	_, err := TicketCollection.InsertOne(ctx, ticket)
	if err != nil {
		log.Println(err)
	}
}

// CheckTicketExists func
func CheckTicketExists(name string) bool {
	if err := TicketCollection.FindOne(ctx, bson.M{"name": name}); err != nil {
		return false
	}
	return true
}
