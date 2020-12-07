package models

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Tickets struct
type Tickets struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	CreatedBy string             `bson:"lastname,omitempty"`
	BugType   string             `bson:"password,omitempty"`
	Status    string             `bson:"status,omitempty"`
	CreatedAt string             `bson:"createdAt,omitempty"`
	UpdatedAt string             `bson:"updatedAt,omitempty"`
}

// GetTicketID func
func (ticket *Tickets) GetTicketID() string {
	return ticket.ID.String()
}

// GetTicketName func
func (ticket *Tickets) GetTicketName() string {
	return ticket.Name
}

// GetTicketCreatedBy func
func (ticket *Tickets) GetTicketCreatedBy() string {
	return ticket.CreatedBy
}

// GetTicketBugType func
func (ticket *Tickets) GetTicketBugType() string {
	return ticket.BugType
}

// GetTicketStatus func
func (ticket *Tickets) GetTicketStatus() string {
	return ticket.Status
}

// GetTicketUpdatedAt func
func (ticket *Tickets) GetTicketUpdatedAt() string {
	return ticket.UpdatedAt
}

// GetTicketCreatedAt func
func (ticket *Tickets) GetTicketCreatedAt() string {
	return ticket.CreatedAt
}

func test() {
	log.Println()
}
