package routes

import (
	"context"
	"net/http"

	"../apperrors"
	"../appsessions"
	"../middleware"
	"../models"
	"../utils"

	"github.com/gorilla/mux"
)

var ctx = context.TODO()

// NewRouter func
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", indexGETHandler).Methods("GET")

	router.HandleFunc("/login", LoginGETHandler).Methods("GET")
	router.HandleFunc("/registration", RegistrationGETHandler).Methods("GET")
	router.HandleFunc("/bugtype/create", BugtypeGETHandler).Methods("GET")
	router.HandleFunc("/ticket/create", TicketsGETHandler).Methods("GET")
	router.HandleFunc("/logout", LogoutGETHandler).Methods("GET")
	router.HandleFunc("/verification", VerificationGETHandler).Methods("GET")
	router.HandleFunc("/profile/{user}", middleware.AuthRequired(ProfileGETHandler)).Methods("GET")
	router.HandleFunc("/ticket/details/{id}", middleware.AuthRequired(TicketDetailGETHandler)).Methods("GET")

	router.HandleFunc("/login", LoginPOSTHandler).Methods("POST")
	router.HandleFunc("/registration", RegistrationPOSTHandler).Methods("POST")
	router.HandleFunc("/bugtype/create", BugtypePOSTHandler).Methods("POST")
	router.HandleFunc("/ticket/create", TicketsPOSTHandler).Methods("POST")
	router.HandleFunc("/verification", VerificationPOSTHandler).Methods("POST")

	fs := http.FileServer(http.Dir("static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	router.NotFoundHandler = router.NewRoute().HandlerFunc(pageNotFoundHandler).GetHandler()

	return router
}

func indexGETHandler(w http.ResponseWriter, r *http.Request) {
	sessionkey := CheckCurrentSession(r)
	if sessionkey == "" {
		utils.ExecuteTemplate(w, "index.html", nil)
		return
	}
	route := "/profile/" + sessionkey
	http.Redirect(w, r, route, 302)

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
	session, _ := appsessions.Store.Get(r, "session")
	key := session.Values["username"]
	if key == "" || key == nil {
		return ""
	}
	exists := models.UserExists(key.(string))
	if exists {
		return key.(string)
	}
	return ""
}

// SaveCurrentSession func
func SaveCurrentSession(w http.ResponseWriter, r *http.Request, key string) error {
	session, _ := appsessions.Store.Get(r, "session")
	session.Values["username"] = key
	err := session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}
