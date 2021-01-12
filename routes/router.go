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

	router.HandleFunc("/", indexGETHandler).Methods("GET")

	router.HandleFunc("/login", loginGETHandler).Methods("GET")
	router.HandleFunc("/registration", registrationGETHandler).Methods("GET")
	router.HandleFunc("/bugtype/create", bugtypeGETHandler).Methods("GET")
	router.HandleFunc("/ticket/create", ticketsGETHandler).Methods("GET")
	router.HandleFunc("/logout", logoutGETHandler).Methods("GET")
	router.HandleFunc("/profile/{user}", middleware.AuthRequired(profileGETHandler)).Methods("GET")

	router.HandleFunc("/login", loginPOSTHandler).Methods("POST")
	router.HandleFunc("/registration", registrationPOSTHandler).Methods("POST")
	router.HandleFunc("/bugtype/create", bugtypePOSTHandler).Methods("POST")
	router.HandleFunc("/ticket/create", ticketsPOSTHandler).Methods("POST")

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

	session, _ := sessions.Store.Get(r, "session")
	session.Values["username"] = username
	session.Save(r, w)

	redirectString := "/profile/" + username

	http.Redirect(w, r, redirectString, 302)
}

func logoutGETHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	delete(session.Values, "username")
	session.Save(r, w)
	http.Redirect(w, r, "/login", 302)
}

func bugtypeGETHandler(w http.ResponseWriter, r *http.Request) {
	user := CheckCurrentSession(r)
	utils.ExecuteTemplate(w, "bugtypes.html", struct {
		User string
	}{
		User: user,
	})
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

	ok := models.BugTypeExists(bugtypeAcronym)

	if ok {
		utils.ExecuteTemplate(w, "bugtypes.html", ok)
	}

	var newBugType models.BugTypes
	newBugType.Acronym = bugtypeAcronym
	newBugType.Name = bugtypeName
	newBugType.Description = bugtypeDescription

	err := newBugType.CreateNewBugType()
	if err != nil {
		utils.ExecuteTemplate(w, "bugtypes.html", err)
	}

	msg := "Bugtype '" + bugtypeAcronym + "' created"

	SaveCurrentSession(w, r, sessionKey)
	utils.ExecuteTemplate(w, "bugtypes.html", msg)
}

func ticketsGETHandler(w http.ResponseWriter, r *http.Request) {
	user := CheckCurrentSession(r)
	list, _ := models.BugTypeListOfAcronyms()

	utils.ExecuteTemplate(w, "tickets.html", struct {
		User string
		List *[]string
	}{
		User: user,
		List: list,
	})
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

func profileGETHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["user"]

	if strings.EqualFold("favicon.ico", username) {
		return
	}
	ok := models.UserExists(username)
	if !ok {
		log.Println(ok)
		return
	}
	user, err := models.UserGetAllInformations(username)
	if err != nil {
		log.Println(err)
	}

	utils.ExecuteTemplate(w, "profile.html", struct {
		User *models.User
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
