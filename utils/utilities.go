package utils

import (
	"log"
	"time"
)

// IsError func
func IsError(err error) {
	if err != nil {
		log.Fatal("An Error occurred : ", err)
	}
}

// CreateTimeStamp func
func CreateTimeStamp() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02 15:04:05")
}
