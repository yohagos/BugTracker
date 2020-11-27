package utils

import (
	"math/rand"
)

// RandomKey func
func RandomKey() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+-!?"
	var length = 30
	var byteSlice []byte
	var result string
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charset))
		randomChar := charset[random]
		byteSlice = append(byteSlice, randomChar)
	}

	result = string(byteSlice)
	return result
}
