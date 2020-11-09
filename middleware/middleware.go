package middleware

import (
	"net/http"

	sess "../sessions"

	"github.com/gorilla/sessions"
)

// HandleFunc type - Handling Authorization
type HandleFunc func(http.ResponseWriter, *http.Request)

// AuthRequired func - checking User Authorization
func AuthRequired(handler HandleFunc) HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := sessions.Store.Get(sess.Store, r, "session")
		_, ok := session.Values["user_id"]
		if !ok {
			http.Redirect(w, r, "/login", 302)
			return
		}
		handler(w, r)
	}
}
