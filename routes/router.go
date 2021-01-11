package routes

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"../apperrors"
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
	router.HandleFunc("/{profile}", middleware.AuthRequired(profilePOSTHandler)).Methods("POST")
	router.HandleFunc("/", indexGETHandler).Methods("GET")

	router.HandleFunc("/registration", registrationGETHandler).Methods("GET")
	router.HandleFunc("/registration", registrationPOSTHandler).Methods("POST")

	router.HandleFunc("/login", loginGETHandler).Methods("GET")
	router.HandleFunc("/login", loginPOSTHandler).Methods("POST")

	router.HandleFunc("/bugtype", bugtypeGETHandler).Methods("GET")
	router.HandleFunc("/bugtype", bugtypePOSTHandler).Methods("POST")

	router.HandleFunc("/ticket", ticketsGETHandler).Methods("GET")
	router.HandleFunc("/ticket", ticketsPOSTHandler).Methods("POST")

	router.HandleFunc("/logout", logoutGETHandler).Methods("GET")

	fs := http.FileServer(http.Dir("static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	router.NotFoundHandler = router.NewRoute().HandlerFunc(pageNotFoundHandler).GetHandler()

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

	err := models.UserAuthentification(username, password)

	if err != nil {
		utils.ExecuteTemplate(w, "login.html", err)
	}

	user, err := models.UserGetAllInformations(username)
	if err != nil {
		log.Println(err)
	}

	session, _ := sessions.Store.Get(r, "session")
	session.Values["username"] = username
	session.Save(r, w)
	username = user.GetUserLastname()
	redirectString := "/" + username

	http.Redirect(w, r, redirectString, 302)
}

func logoutGETHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	delete(session.Values, "username")
	session.Save(r, w)
	http.Redirect(w, r, "/login", 302)
}

func bugtypeGETHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "createbugs.html", nil)
}

func bugtypePOSTHandler(w http.ResponseWriter, r *http.Request) {
	sessionKey := CheckCurrentSession(r)
	if sessionKey == "" {
		utils.ExecuteTemplate(w, "bugtypes.html", apperrors.ErrorSessionInvalid)
		time.Sleep(15 * time.Second)
		http.Redirect(w, r, "/", 303)
	}
	r.ParseForm()

	bugtypeAcronym := r.PostForm.Get("acronym")
	bugtypeName := r.PostForm.Get("name")
	bugtypeDescription := r.PostForm.Get("description")

	ok := models.NewBugTypeExists(bugtypeAcronym)

	if ok != nil {
		utils.ExecuteTemplate(w, "bugtypes.html", ok)
	}

	var newBugType models.BugTypes
	newBugType.Acronym = bugtypeAcronym
	newBugType.Name = bugtypeName
	newBugType.Description = bugtypeDescription

	newBugType.CreateNewBugType()

	SaveCurrentSession(w, r, sessionKey)
	http.Redirect(w, r, "/bugtype", 302)
}

func ticketsGETHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "tickets.html", nil)
}

func ticketsPOSTHandler(w http.ResponseWriter, r *http.Request) {
	sessionKey := CheckCurrentSession(r)
	if sessionKey == "" {
		utils.ExecuteTemplate(w, "tickets.html", apperrors.ErrorSessionInvalid)
		time.Sleep(15 * time.Second)
		http.Redirect(w, r, "/", 303)
	}
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	ok := models.UserAuthentification(username, password)

	if ok != nil {
		utils.ExecuteTemplate(w, "login.gohtml", apperrors.ErrorRoutesInvalidLogin)
	}

	session, _ := sessions.Store.Get(r, "session")
	session.Values["username"] = username
	session.Save(r, w)

	http.Redirect(w, r, "/"+username, 302)
}

func profilePOSTHandler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["profile"]
	
	ok := models.UserExists(username)
	if !ok {
		log.Println(ok)
		http.Redirect(w, r, "/login", 302)
	}
	user, err := models.UserGetAllInformations(currentUser.(string))
	if err != nil {
		log.Println(err)
	}

	utils.ExecuteTemplate(w, "profile.html", struct {
		User models.User
	}{
		User: user,
	})
}

func pageNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "error.html", struct {
		Error error
	}{
		Error: apperrors.ErrorRoutesPageNotFound,
	})
}

// CheckCurrentSession func
func CheckCurrentSession(r *http.Request) string {
	session, _ := sessions.Store.Get(r, "session")
	key := session.Values["username"]
	exists := models.UserExists(key.(string))
	if exists {
		return key.(string)
	}
	return ""
}

// SaveCurrentSession func
func SaveCurrentSession(w http.ResponseWriter, r *http.Request, key string) error {
	session, _ := sessions.Store.Get(r, "session")
	session.Values["username"] = key
	err := session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}
