package routes

import (
	"context"
	"log"
	"net/http"
	"strings"

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

	router.HandleFunc("/registration", registrationGETHandler).Methods("GET")
	router.HandleFunc("/registration", registrationPOSTHandler).Methods("POST")

	router.HandleFunc("/login", loginGETHandler).Methods("GET")
	router.HandleFunc("/login", loginPOSTHandler).Methods("POST")

	router.HandleFunc("/login", bugtypeGETHandler).Methods("GET")
	router.HandleFunc("/login", bugtypePOSTHandler).Methods("POST")

	router.HandleFunc("/login", ticketsGETHandler).Methods("GET")
	router.HandleFunc("/login", ticketsPOSTHandler).Methods("POST")

	router.HandleFunc("/logout", logoutGETHandler).Methods("GET")

	router.HandleFunc("/{profile}", middleware.AuthRequired(profileGETHandler)).Methods("GET")

	fs := http.FileServer(http.Dir("../static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	router.NotFoundHandler = router.NewRoute().HandlerFunc(pageNotFoundHandler).GetHandler()

	return router
}

func indexGETHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "index.gohtml", nil)
}

func registrationGETHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "registration.gohtml", nil)
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
	utils.ExecuteTemplate(w, "login.gohtml", nil)
}

func loginPOSTHandler(w http.ResponseWriter, r *http.Request) {
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

func logoutGETHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	delete(session.Values, "username")
	session.Save(r, w)
	http.Redirect(w, r, "/login", 302)
}

func bugtypeGETHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "createbugs.gohtml", nil)
}

func bugtypePOSTHandler(w http.ResponseWriter, r *http.Request) {
	/* r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	ok := models.UserAuthentification(username, password)

	if ok != nil {
		utils.ExecuteTemplate(w, "login.gohtml", apperrors.ErrorRoutesInvalidLogin)
	}

	session, _ := sessions.Store.Get(r, "session")
	session.Values["username"] = username
	session.Save(r, w)

	http.Redirect(w, r, "/"+username, 302) */
}

func ticketsGETHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "tickets.gohtml", nil)
}

func ticketsPOSTHandler(w http.ResponseWriter, r *http.Request) {
	/* r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	ok := models.UserAuthentification(username, password)

	if ok != nil {
		utils.ExecuteTemplate(w, "login.gohtml", apperrors.ErrorRoutesInvalidLogin)
	}

	session, _ := sessions.Store.Get(r, "session")
	session.Values["username"] = username
	session.Save(r, w)

	http.Redirect(w, r, "/"+username, 302) */
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
	user, err := models.UserGetAllInformations(currentUser.(string))
	if err != nil {
		log.Println(err)
	}

	utils.ExecuteTemplate(w, "profile.gohtml", struct {
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
