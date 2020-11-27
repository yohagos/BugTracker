package sessions

import (
	"../utils"
	"github.com/gorilla/sessions"
)

// Store var - CookieStore
var (
	Store *sessions.CookieStore
)

// SessionInit func - initialize Session
func SessionInit() {
	authKeyOne := utils.RandomKey()

	Store = sessions.NewCookieStore([]byte(authKeyOne))

	Store.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		MaxAge:   60 * 120}
}
