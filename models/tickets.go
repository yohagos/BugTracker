package models

import (
	"fmt"
	"log"

	"../databases"
	"../utils"

	"go.mongodb.org/mongo-driver/bson"
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

// TestCreateTicket func
func TestCreateTicket() {
	timestamp := utils.CreateTimeStamp()
	ticketDocument := bson.D{
		{Key: "name", Value: "test-12345"},
		{Key: "createdby", Value: "test01"},
		{Key: "bugtype", Value: "bug"},
		{Key: "status", Value: "testing"},
		{Key: "createdAt", Value: timestamp},
		{Key: "updatedAt", Value: timestamp},
	}

	databases.CreateNewTicket(ticketDocument)
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

func (ticket *Tickets) setTicketName(tic string) {
	ticket.Name = tic
}

func (ticket *Tickets) setTicketCreatedBy(tic string) {
	ticket.CreatedBy = tic
}

func (ticket *Tickets) setTicketBugType(tic string) {
	ticket.BugType = tic
}

func (ticket *Tickets) setTicketStatus(tic string) {
	ticket.Status = tic
}

func (ticket *Tickets) setTicketUpdatedAt(tic string) {
	ticket.UpdatedAt = tic
}

func (ticket *Tickets) setTicketCreatedAt(tic string) {
	ticket.CreatedAt = tic
}

// CreateNewTicket func
func (ticket *Tickets) CreateNewTicket() {
	ok := databases.CheckTicketExists(ticket.GetTicketName())
	if !ok {
		log.Println("Ticket already exists")
		return
	}
	num := utils.RandomFiveDigitNumber()
	timeStamp := utils.CreateTimeStamp()
	ticketDocument := bson.D{
		{Key: "name", Value: ticket.GetTicketName() + "-" + num},
		{Key: "createdby", Value: ticket.GetTicketCreatedBy()},
		{Key: "bugtype", Value: ticket.GetTicketBugType()},
		{Key: "status", Value: ticket.GetTicketStatus()},
		{Key: "createdAt", Value: timeStamp},
		{Key: "updatedAt", Value: timeStamp},
	}
	databases.CreateNewTicket(ticketDocument)
}

// NewTicketExists func
func NewTicketExists(name string) bool {
	return databases.CheckTicketExists(name)
}

// TicketGetAllInformations func
func TicketGetAllInformations(name string) (Tickets, error) {
	var ticket Tickets

	result, err := databases.GetAllTicketInformations(name)
	if err != nil {
		log.Println(err)
		return ticket, err
	}

	for k, v := range result {
		switch k {
		case "name":
			key := fmt.Sprintf("%v", v)
			ticket.setTicketName(key)
		case "bugtype":
			key := fmt.Sprintf("%v", v)
			ticket.setTicketBugType(key)
		case "createdBy":
			key := fmt.Sprintf("%v", v)
			ticket.setTicketCreatedBy(key)
		case "status":
			key := fmt.Sprintf("%v", v)
			ticket.setTicketStatus(key)
		case "createdAt":
			key := fmt.Sprintf("%v", v)
			ticket.setTicketCreatedAt(key)
		case "updatedAt":
			key := fmt.Sprintf("%v", v)
			ticket.setTicketUpdatedAt(key)
		default:

		}
	}

	return ticket, nil
}
