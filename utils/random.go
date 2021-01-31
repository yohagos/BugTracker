package utils

import (
	"math/rand"
	"strconv"
	"time"
)

// CHARSET const
const CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+-!?"

// RandomKey func
func RandomKey() string {
	return string(randStringBytes(30))
}

// RandomFiveDigitNumber func
func RandomFiveDigitNumber() string {
	rand.Seed(time.Now().UnixNano())
	min := 10000
	max := 99999
	randomInt := rand.Intn(max-min) + min
	return strconv.Itoa(randomInt)
}

// GenerateVerificationKey func
func GenerateVerificationKey() string {
	return string(randStringBytes(6))
}

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = CHARSET[rand.Intn(len(CHARSET))]
	}
	return string(b)
}
