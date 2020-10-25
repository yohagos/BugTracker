package utils

import (
	"log"
)

func IsError(err error) {
	if err != nil {
		log.Fatal("An Error accured : ", err)
	}
}
