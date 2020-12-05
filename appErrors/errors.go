package utils

import (
	"errors"
	"net/http"
)

var (

	// ErrorRoutesPageNotFound var
	ErrorRoutesPageNotFound = errors.New("Page does not exist. Please try again")
	// ErrorRoutesPageDoesntExist var
	ErrorRoutesPageDoesntExist = errors.New("This page does not exist")
	// ErrorRoutesUserNotFound var
	ErrorRoutesUserNotFound = errors.New("Username not found")
	// ErrorRoutesInvalidLogin var
	ErrorRoutesInvalidLogin = errors.New("Invalid login. Please try again")

	// ErrorUserDoesNotExist var
	ErrorUserDoesNotExist = errors.New("Login is invalid. Username / Password does not exists")

	// ErrorBugTypeAlreadyExists func
	ErrorBugTypeAlreadyExists = errors.New("BugType already exists")
)

// InternalServerError func
func InternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal Server Error"))
}
