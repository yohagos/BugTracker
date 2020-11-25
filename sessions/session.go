package sessions

import (
	"github.com/gorilla/sessions"
	/* "github.com/kidstuff/mongostore"

	"github.com/globalsign/mgo" */)

/*  */

/* const (
	key = "xT1VnWhYCgA32lIYZfPMPaGkj3XCyj"
) */

/* var (
	// DBSession var
	DBSession *mgo
) */

/* func SessionInit() {
	DBSession, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	//defer DBSession.Close()

	store := mongostore.NewMongoStore(dbSession.DB("sessions").C("test_session"), 3600, true, []byte(key))

	session, err := store.Get(request, "username")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(session)
} */

const (
	key = "xT1VnWhYCgA32lIYZfPMPaGkj3XCyj"
)

// Store var - CookieStore
var Store = sessions.NewCookieStore([]byte(key))

// SessionInit func - initialize Session
func SessionInit() {
	Store.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		MaxAge:   60 * 120}
}

// GetSessionStore func
func GetSessionStore() *sessions.CookieStore {
	return Store
}
