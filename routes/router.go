package routes

import (
	"context"
	"log"
	"net/http"

	"../databases"
	"../middleware"
	"../models"
	"../sessions"
	"../utils"

	"github.com/gorilla/mux"
)

var ctx = context.TODO()

// NewRouter func
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", middleware.AuthRequired(indexGETHandler)).Methods("GET")
	router.HandleFunc("/test", testGETHandler).Methods("GET")
	router.HandleFunc("/registration", registrationGETHandler).Methods("GET")
	router.HandleFunc("/registration", registrationPOSTHandler).Methods("POST")

	router.HandleFunc("/login", loginGETHandler).Methods("GET")
	router.HandleFunc("/login", loginPOSTHandler).Methods("POST")

	fs := http.FileServer(http.Dir("../static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	return router
}

func indexGETHandler(w http.ResponseWriter, r *http.Request) {
	/* session, _ := sessions.Store.Get(r, "session")
	if err != nil {
		utils.InternalServerError(w)
	}
	//ok = databases.UserExists(currentUser.(string))
	currentUser, ok := session.Values["username"]
	if !ok {
		log.Println(ok)
		http.Redirect(w, r, "/login", 302)
	}
	username, ok := currentUser.(string)
	if !ok {
		log.Println(ok)
		http.Redirect(w, r, "/login", 302)
	} */

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

	/* ok := databases.UserAuthentification(username, password)

	if ok != nil {
		utils.ExecuteTemplate(w, "login.html", utils.ErrorInvalidLogin)
	} */

	log.Println(username + " " + password)

	session, err := sessions.Store.Get(r, "session")
	utils.IsError(err)
	session.Values["username"] = username

	session.Save(r, w)
	//utils.IsError(err)

	http.Redirect(w, r, "/test", 302)
}

func testGETHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	currentUser, ok := session.Values["username"]
	if !ok {
		log.Println(ok)
		http.Redirect(w, r, "/login", 302)
	}
	username, ok := currentUser.(string)
	if !ok {
		log.Println(ok)
		http.Redirect(w, r, "/login", 302)
	}

	w.Write([]byte(username))
}
