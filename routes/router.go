package routes

import (
	"context"
	"log"
	"net/http"

	"../databases"
	"../models"
	"../utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

var ctx = context.TODO()
var newuser *models.User

// NewRouter func
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", indexGETHandler).Methods("GET")

	router.HandleFunc("/registration", registrationGETHandler).Methods("GET")
	router.HandleFunc("/registration", registrationPOSTHandler).Methods("POST")

	fs := http.FileServer(http.Dir("../static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	return router
}

func indexGETHandler(w http.ResponseWriter, r *http.Request) {
	Users, err := databases.GetAllUsers()
	if err != nil {
		log.Fatal(err)
	}
	if len(Users) <= 0 {
		utils.ExecuteTemplate(w, "index.html", nil)
	} else {
		utils.ExecuteTemplate(w, "index.html", struct {
			User []bson.M
		}{
			User: Users,
		})
	}

}

/* func indexPOSTHandler(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/", 302)
} */

func registrationGETHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "registration.html", nil)
}

func registrationPOSTHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userName := r.PostForm.Get("name")
	userLastname := r.PostForm.Get("lastname")
	userEmail := r.PostForm.Get("email")
	userPassword := r.PostForm.Get("password")

	newuser = models.CreateNewUser(userName, userLastname, userEmail, userPassword)
	databases.CreateUser(newuser)
	http.Redirect(w, r, "/", 302)
}
