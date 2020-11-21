package routes

import (
	"context"
	"fmt"
	"net/http"

	"../databases"
	"../models"
	"../utils"

	"github.com/gorilla/mux"
)

var ctx = context.TODO()
var newuser *models.User

// NewRouter func
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	// router.HandleFunc("/", middleware.AuthRequired(indexGETHandler)).Methods("GET")
	router.HandleFunc("/", indexGETHandler).Methods("GET")

	router.HandleFunc("/registration", registrationGETHandler).Methods("GET")
	router.HandleFunc("/registration", registrationPOSTHandler).Methods("POST")

	router.HandleFunc("/login", loginGETHandler).Methods("GET")
	router.HandleFunc("/login", loginPOSTHandler).Methods("POST")

	fs := http.FileServer(http.Dir("../static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	return router
}

func indexGETHandler(w http.ResponseWriter, r *http.Request) {
	/* session, _ := sessions.Store.Get(sess.Store, r, "session")
	untypeduser_id := session.Values["user_id"]
	currentUser, _ := untypeduser_id.(int64)

	fmt.Println(currentUser) */
	utils.ExecuteTemplate(w, "index.html", nil)
}

func registrationGETHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "registration.html", nil)
}

func registrationPOSTHandler(w http.ResponseWriter, r *http.Request) {
	var createUser models.User
	r.ParseForm()
	createUser.Name = r.PostForm.Get("name")
	createUser.Lastname = r.PostForm.Get("lastname")
	createUser.Email = r.PostForm.Get("email")
	createUser.Password = r.PostForm.Get("password")

	databases.AddNewUser(models.CreateNewUser(createUser))

	http.Redirect(w, r, "/login", 302)
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
