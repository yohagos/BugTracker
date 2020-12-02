package routes

import (
	"context"
	"log"
	"net/http"
	"strings"

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
	router.HandleFunc("/", indexGETHandler).Methods("GET")

	router.HandleFunc("/registration", registrationGETHandler).Methods("GET")
	router.HandleFunc("/registration", registrationPOSTHandler).Methods("POST")

	router.HandleFunc("/login", loginGETHandler).Methods("GET")
	router.HandleFunc("/login", loginPOSTHandler).Methods("POST")

	router.HandleFunc("/logout", logoutGETHandler).Methods("GET")

	router.HandleFunc("/{profile}", middleware.AuthRequired(profileGETHandler)).Methods("GET")

	fs := http.FileServer(http.Dir("../static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	return router
}

func indexGETHandler(w http.ResponseWriter, r *http.Request) {
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

	createUser.CreateNewUser()

	http.Redirect(w, r, "/login", 302)
}

func loginGETHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

func loginPOSTHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	ok := models.UserAuthentification(username, password)

	if ok != nil {
		utils.ExecuteTemplate(w, "login.html", utils.ErrorInvalidLogin)
	}

	session, _ := sessions.Store.Get(r, "session")
	session.Values["username"] = username
	session.Save(r, w)

	http.Redirect(w, r, "/"+username, 302)
}

func profileGETHandler(w http.ResponseWriter, r *http.Request) {
	i := r.URL.RequestURI()[1:]
	if strings.EqualFold("favicon.ico", i) {
		return
	}

	session, _ := sessions.Store.Get(r, "session")
	currentUser, ok := session.Values["username"]
	if !ok {
		log.Println(ok)
		http.Redirect(w, r, "/login", 302)
	}

	ok = models.UserExists(currentUser.(string))
	if !ok {
		log.Println(ok)
		http.Redirect(w, r, "/login", 302)
	}
	user := models.GetUserInformations(currentUser.(string))

	utils.ExecuteTemplate(w, "profile.html", struct {
		User *models.User
	}{
		User: user,
	})
}

func logoutGETHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	delete(session.Values, "username")
	session.Save(r, w)
	http.Redirect(w, r, "/login", 302)
}
