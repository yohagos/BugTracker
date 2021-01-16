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
		return true
	}
	return false
}

// GetAllTicketInformations func
func GetAllTicketInformations(name string) (bson.M, error) {
	var result bson.M
	if err := TicketCollection.FindOne(ctx, bson.M{"name": name}).Decode(&result); err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

// AllTicketsByUser func
func AllTicketsByUser(username string) (*[]bson.M, error) {
	cursor, err := TicketCollection.Find(ctx, bson.M{"createdby": username})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var ticketList []bson.M
	if err = cursor.All(ctx, &ticketList); err != nil {
		log.Println(err)
		return nil, err
	}

	return &ticketList, nil
}
