package apperrors

import (
	"errors"
	"net/http"
)

var (
	// ErrorSessionInvalid error
	ErrorSessionInvalid = errors.New("Current Session is invalid. Please login first")

	// ErrorRoutesPageNotFound error
	ErrorRoutesPageNotFound = errors.New("Page does not exist. Please try again")
	// ErrorRoutesPageDoesntExist error
	ErrorRoutesPageDoesntExist = errors.New("This page does not exist")
	// ErrorRoutesUserNotFound error
	ErrorRoutesUserNotFound = errors.New("Username not found")
	// ErrorRoutesInvalidLogin error
	ErrorRoutesInvalidLogin = errors.New("Invalid login. Please try again")

	// ErrorUserDoesNotExist error
	ErrorUserDoesNotExist = errors.New("Login is invalid. Username / Password does not exists")

	// ErrorBugTypeAlreadyExists error
	ErrorBugTypeAlreadyExists = errors.New("BugType already exists")

	// ErrorTicketAlreadyExits error
	ErrorTicketAlreadyExits = errors.New("Ticket already exits")
)

// InternalServerError func
func InternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal Server Error"))
}
