package utils

import (
	"fmt"
	"math/rand"
)

func RandomKey() {
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
	fmt.Println(result)
}
