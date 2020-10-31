package routes

import (
	"context"
	"net/http"

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
	/* var ok bool
	ok = false
	if len(newuser.GetUserName()) > 0 {
		ok = true
	} */
	utils.ExecuteTemplate(w, "index.html", struct {
		/* Display bool */
		User *models.User
	}{
		/* Display: ok, */
		User: newuser,
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
	http.Redirect(w, r, "/", 302)
}
