package routes

import (
	"log"
	"net/http"
	"strings"

	"../appsessions"
	"../mails"
	"../models"
	"../utils"
	"github.com/gorilla/mux"
)

// RegistrationGETHandler func
func RegistrationGETHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "registration.html", nil)
}

// RegistrationPOSTHandler func
func RegistrationPOSTHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	key := utils.GenerateVerificationKey()
	mail := r.PostForm.Get("email")
	name := r.PostForm.Get("name")

	var verificationUser models.UserVerification

	verificationUser.SetUserVerificationName(name)
	verificationUser.SetUserVerificationLastname(r.PostForm.Get("lastname"))
	verificationUser.SetUserVerificationPassword(r.PostForm.Get("password"))
	verificationUser.SetUserVerificationEmail(mail)
	verificationUser.SetUserVerificationGeneratedKey(key)

	verificationUser.CreateVerificationProfile()

	mails.SendVerificationMail(name, mail, key)

	http.Redirect(w, r, "/verification", 302)
}

// LoginGETHandler func
func LoginGETHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

// LoginPOSTHandler func
func LoginPOSTHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	err := models.UserAuthentification(username, password)

	if err != nil {
		utils.ExecuteTemplate(w, "login.html", err)
	}

	SaveCurrentSession(w, r, username)

	redirectString := "/profile/" + username

	http.Redirect(w, r, redirectString, 302)
}

// LogoutGETHandler func
func LogoutGETHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := appsessions.Store.Get(r, "session")
	delete(session.Values, "username")
	session.Save(r, w)
	http.Redirect(w, r, "/login", 302)
}

// ProfileGETHandler func
func ProfileGETHandler(w http.ResponseWriter, r *http.Request) {
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

	ticketsList, err := models.GetTicketsByUser(username)
	if err != nil {
		log.Println(err)
	}

	utils.ExecuteTemplate(w, "profile.html", struct {
		User    *models.User
		Tickets *[]models.Tickets
	}{
		User:    user,
		Tickets: ticketsList,
	})
}
