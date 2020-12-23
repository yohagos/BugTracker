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
