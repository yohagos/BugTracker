package routes

import (
	"log"
	"net/http"
	"strings"

	"../mails"
	"../models"
	"../sessions"
	"../utils"
	"github.com/gorilla/mux"
)

// RegistrationGETHandler func
func RegistrationGETHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "registration.html", nil)
}

// RegistrationPOSTHandler func
func RegistrationPOSTHandler(w http.ResponseWriter, r *http.Request) {
	/* var createUser models.User
	r.ParseForm()
	createUser.Name = r.PostForm.Get("name")
	createUser.Lastname = r.PostForm.Get("lastname")
	createUser.Email = r.PostForm.Get("email")
	createUser.Password = r.PostForm.Get("password")

	createUser.CreateNewUser() */
	key := utils.GenerateVerificationKey()
	mail := r.PostForm.Get("email")
	name := r.PostForm.Get("name")

	var verificationUser models.UserVerification

	r.ParseForm()
	verificationUser.SetUserVerificationName(name)
	verificationUser.SetUserVerificationLastname(r.PostForm.Get("lastname"))
	verificationUser.SetUserVerificationPassword(r.PostForm.Get("password"))
	verificationUser.SetUserVerificationEmail(mail)

	verificationUser.SetUserVerificationGeneratedKey(key)
	/* verificationUser.SetUserVerificationVerified(false) */

	verificationUser.CreateVerificationProfile()

	mails.SendVerificationMail(name, mail, key)

	/* route := "/verfication/" + mail */

	http.Redirect(w, r, "/verfication", 302)
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

	session, _ := sessions.Store.Get(r, "session")
	session.Values["username"] = username
	session.Save(r, w)

	redirectString := "/profile/" + username

	http.Redirect(w, r, redirectString, 302)
}

// LogoutGETHandler func
func LogoutGETHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
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
