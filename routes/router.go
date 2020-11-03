package routes

import (
	"context"
	"log"
	"net/http"

	"../database"
	"../models"
	"../utils"
	"github.com/gorilla/mux"
)

var ctx = context.TODO()
var newuser *models.User

// NewRouter func
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", indexGETHandler).Methods("GET")
	router.HandleFunc("/", indexPOSTHandler).Methods("POST")

	fs := http.FileServer(http.Dir("../static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	return router
}

func indexGETHandler(w http.ResponseWriter, r *http.Request) {
	var display bool
	display = false
	/* if len(newuser.GetUserName()) > 0 {
		display = true
	} */
	utils.ExecuteTemplate(w, "index.html", struct {
		Display bool
		User    *models.User
	}{
		Display: display,
		User:    newuser,
	})
}

func indexPOSTHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userName := r.PostForm.Get("name")
	userLastname := r.PostForm.Get("lastname")
	userEmail := r.PostForm.Get("email")
	userPassword := r.PostForm.Get("password")

	//models.users.SetA{userName, userLastname, userEmail, userPassword}
	newuser = models.CreateNewUser(userName, userLastname, userEmail, userPassword)
	log.Println(newuser)
	database.CreateUser(newuser)
	http.Redirect(w, r, "/", 302)
}
