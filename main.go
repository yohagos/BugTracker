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
	os.Exit(3) */

	sessions.SessionInit()

	log.Println("Loading Templates..")
	utils.LoadTemplate("static/*.html")

	r := routes.NewRouter()
	log.Println("Server starting..")
	if err := http.ListenAndServe(":8888", r); err != nil {
		log.Fatal("Server error! Message : ", err)
	}
}
