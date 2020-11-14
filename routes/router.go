package routes

import (
	"context"
	"fmt"
	"net/http"

	"../middleware"
	"../models"
	sess "../sessions"
	"../utils"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var ctx = context.TODO()
var newuser *models.User

// NewRouter func
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", middleware.AuthRequired(indexGETHandler)).Methods("GET")

	router.HandleFunc("/registration", registrationGETHandler).Methods("GET")
	router.HandleFunc("/registration", registrationPOSTHandler).Methods("POST")

	router.HandleFunc("/login", loginGETHandler).Methods("GET")
	router.HandleFunc("/login", loginPOSTHandler).Methods("POST")

	fs := http.FileServer(http.Dir("../static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	return router
}

func indexGETHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(sess.Store, r, "session")
	untypeduser_id := session.Values["user_id"]
	currentUser, _ := untypeduser_id.(int64)

	fmt.Println(currentUser)
	utils.ExecuteTemplate(w, "index.html", nil)
}

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
	//databases.CreateUser(newuser)
	http.Redirect(w, r, "/", 302)
}

func loginGETHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

func loginPOSTHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	fmt.Println(username + password)
}
