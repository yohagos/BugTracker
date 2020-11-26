package sessions

import (
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

// Store var - CookieStore
var Store *sessions.CookieStore

// SessionInit func - initialize Session
func SessionInit() {
	authKeyOne := securecookie.GenerateRandomKey(64)
	//encryptedKeyOne := securecookie.GenerateRandomKey(32)

	/* Store = sessions.NewCookieStore(
		authKeyOne,
		//encryptedKeyOne,
	) */

	Store = sessions.NewCookieStore(
		[]byte(authKeyOne),
	)

	Store.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		MaxAge:   60 * 120}
}

/* // GetSessionStore func
func GetSessionStore() *sessions.CookieStore {
	return Store
} */
