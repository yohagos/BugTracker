package main

import (
	"context"
	"log"
	"net/http"

	"./database"
	"./routes"
	utils "./utils"
)

var ctx = context.TODO()

func main() {
	log.Println("Initilate Database..")
	database.Init()

	/* var list []string = {"Yosef", "Hagos", "test@test.de", "test123"}
	models.CreateNewUser(list) */

	log.Println("Loading Templates..")
	utils.LoadTemplate("static/*.html")

	r := routes.NewRouter()
	log.Println("Server starting..")
	if err := http.ListenAndServe(":8888", r); err != nil {
		log.Fatal("Server error! Message : ", err)
	}
}

func ifError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
