package utils

import (
	"log"
)

// IsError func
func IsError(err error) {
	if err != nil {
		log.Fatal("An Error occurred : ", err)
	}
}
