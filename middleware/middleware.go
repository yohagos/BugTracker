package middleware

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type HandleFunc func(http.ResponseWriter, *http.Request)

func AuthRequired(handler HandleFunc) HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := sessions.Store.Get(r, "session")
		_, ok := session.Values["user_id"]
		if !ok {
			http.Redirect(w, r, "/login", 302)
			return
		}
		handler(w, r)
	}
}
