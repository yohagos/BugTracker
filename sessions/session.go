package sessions

import (
	"github.com/gorilla/sessions"
)

const (
	key = "xT1VnWhYCgA32lIYZfPMPaGkj3XCyj"
)

// Store var - CookieStore
var Store = sessions.NewCookieStore([]byte(key))

// SessionInit func - initialize Session
func SessionInit() {
	Store.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 60 * 120}
}

// GetSessionStore func
func GetSessionStore() *sessions.CookieStore {
	return Store
}
