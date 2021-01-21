package main

import (
	"log"
	"net/http"
	"os"

	"./databases"
	"./mails"
	"./routes"
	"./sessions"
	"./utils"
)

func main() {
	log.Println("Initilate Database..")
	databases.Init()
	mails.SMTPServer("Yosie", "bugtracker2021@gmail.com", "123456")
	/* mails.SendMail() */
	//mails.SendVerificationMail("Yosie", "bugtracker2021@gmail.com", "abcdef")
	os.Exit(4)

	/* models.TestCreateUser()
	models.TestCreateTicket()
	models.TestCreateNewBugType() */
	/* os.Exit(3) */

	sessions.SessionInit()

	log.Println("Loading Templates..")
	utils.LoadTemplate("./templates/*.html")

	r := routes.NewRouter()

	log.Println("Server starting..")
	if err := http.ListenAndServe(":8888", r); err != nil {
		log.Fatal("Server error! Message : ", err)
	}
}
