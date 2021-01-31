package routes

import (
	"log"
	"net/http"
	"time"

	"../apperrors"
	"../models"
	"../utils"
	"github.com/gorilla/mux"
)

// TicketsGETHandler func
func TicketsGETHandler(w http.ResponseWriter, r *http.Request) {
	user := CheckCurrentSession(r)
	list, _ := models.BugTypeListOfAcronyms()

	utils.ExecuteTemplate(w, "tickets.html", struct {
		User string
		List []string
	}{
		User: user,
		List: list,
	})
}

// TicketsPOSTHandler func
func TicketsPOSTHandler(w http.ResponseWriter, r *http.Request) {
	sessionKey := CheckCurrentSession(r)
	if sessionKey == "" {
		utils.ExecuteTemplate(w, "tickets.html", apperrors.ErrorSessionInvalid)
		time.Sleep(15 * time.Second)
		http.Redirect(w, r, "/", 303)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		bugtype := r.PostForm.Get("bugtype")
		status := r.PostForm.Get("status")

		var ticket models.Tickets
		ticket.SetTicketBugType(bugtype)
		ticket.SetTicketStatus(status)
		ticket.SetTicketCreatedBy(sessionKey)

		ticket.CreateNewTicket()

		SaveCurrentSession(w, r, sessionKey)

		redirect := "/profile/" + sessionKey

		http.Redirect(w, r, redirect, 302)
	}
}

// TicketDetailGETHandler func
func TicketDetailGETHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	ticket, err := models.TicketGetAllInformations(id)
	if err != nil {
		log.Println(err)
	}

	user, err := models.UserGetAllInformations(ticket.GetTicketCreatedBy())
	if err != nil {
		log.Println(err)
	}

	utils.ExecuteTemplate(w, "ticketdetails.html", struct {
		Ticket *models.Tickets
		User   *models.User
	}{
		Ticket: ticket,
		User:   user,
	})
}
