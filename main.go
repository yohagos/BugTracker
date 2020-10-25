package main

import (
	"context"
	"log"
	"net/http"

	"./routes"
	utils "./utils"
)

var (
	msgList []string
)

var ctx = context.TODO()

func main() {
	utils.LoadTemplate("static/*.html")

	r := routes.NewRouter()
	routes.SetList(msgList)
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
