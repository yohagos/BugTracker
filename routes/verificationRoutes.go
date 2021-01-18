package routes

import (
	"errors"
	"net/http"

	"../models"
	"../utils"
	"github.com/gorilla/mux"
)

// VerificationGETHandler func
func VerificationGETHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "verification.html", nil)
}

// VerificationPOSTHandler func
func VerificationPOSTHandler(w http.ResponseWriter, r *http.Request) {
	user := mux.Vars(r)["user"]
	r.ParseForm()
	input := r.PostForm.Get("verification")

	if r.Method == http.MethodPost {
		bo := models.CheckVerification(user, input)

		if !bo {
			utils.ExecuteTemplate(w, "verfication.html", struct {
				Error error
			}{
				Error: errors.New("Wrong key - please try again"),
			})
		}

		models.CreateNewUser(user)
	}
	http.Redirect(w, r, "/login", 302)

}
