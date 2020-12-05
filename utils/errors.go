package utils

import (
	"errors"
	"net/http"
)

var (
	// ErrorPageNotFound var
	ErrorPageNotFound = errors.New("Page does not exist. Please try again")
	// ErrorPageDoesntExist var
	ErrorPageDoesntExist = errors.New("This page does not exist")
	// ErrorUserNotFound var
	ErrorUserNotFound = errors.New("Username not found")
	// ErrorInvalidLogin var
	ErrorInvalidLogin = errors.New("Invalid login. Please try again")
	// ErrorUserDoesNotExist var
	ErrorUserDoesNotExist = errors.New("Login is invalid. Username / Password does not exists")
)

// InternalServerError func
func InternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal Server Error"))
}
