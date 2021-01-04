package main

import (
	"log"
	"net/http"

	"./databases"
	"./routes"
	"./sessions"
	"./utils"
)

func main() {
	log.Println("Initilate Database..")
	databases.Init()

	/* models.TestCreateUser()
	models.TestCreateNewBugType()
	models.TestCreateTicket()

	log.Println(models.UserGetAllInformations("test@test.com"))
	log.Println(models.TicketGetAllInformations("test-12345"))
	log.Println(models.BugTypeGetAllInformations("BUG"))
	os.Exit(3) */

	sessions.SessionInit()

	log.Println("Loading Templates..")
	utils.LoadTemplate("./templates/*.html")

	r := routes.NewRouter()

	log.Println("Server starting..")
	if err := http.ListenAndServe(":8888", r); err != nil {
		log.Fatal("Server error! Message : ", err)
	}
}
