package routes

import (
	"context"
	"net/http"

	utils "../utils"

	"github.com/gorilla/mux"
)

var ctx = context.TODO()

var list []string

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", indexGETHandler).Methods("GET")
	router.HandleFunc("/", indexPOSTHandler).Methods("POST")

	fs := http.FileServer(http.Dir("../static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	return router
}

func indexGETHandler(w http.ResponseWriter, r *http.Request) {
	var displayList bool
	displayList = false
	if len(list) > 0 {
		displayList = true
	}

	utils.ExecuteTemplate(w, "index.html", struct {
		Display  bool
		Messages []string
	}{
		Display:  displayList,
		Messages: list,
	})
}

func indexPOSTHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	message := r.PostForm.Get("txtArea")
	list = append(list, message)
	http.Redirect(w, r, "/", 302)
}

func SetList(msgList []string) {
	list = msgList
}
