package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"./databases"
	"./routes"
	utils "./utils"
)

var ctx = context.TODO()

func main() {

	//os.Exit(0)
	log.Println("Initilate Database..")
	databases.Init()
	/* databases.FindDocument("Basir")
	databases.UpdateUser("Yosie") */
	/* databases.FindDocument("Yosie")
	databases.UpdateUser("Basir") */
	os.Exit(0)

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
