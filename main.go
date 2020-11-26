package main

import (
	"context"
	"log"
	"net/http"

	"./databases"
	"./routes"
	utils "./utils"
)

var ctx = context.TODO()

func main() {
	log.Println("Initilate Database..")
	databases.Init()
	//databases.GetUserInformations()
	//databases.TestUser()
	//os.Exit(3)

	log.Println("Loading Templates..")
	utils.LoadTemplate("static/*.html")

	r := routes.NewRouter()
	log.Println("Server starting..")
	if err := http.ListenAndServe(":8888", r); err != nil {
		log.Fatal("Server error! Message : ", err)
	}
}
