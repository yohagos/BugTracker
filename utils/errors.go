package utils

import "errors"

var (
	ErrorPageNotFound    = errors.New("Page does not exist. Please try again")
	ErrorPageDoesntExist = errors.New("This page does not exist.")
	ErrorUserNotFOund    = errors.New("Username not found")
	ErrorInvalidLogin    = errors.New("Invalid login. Please try again")
)
