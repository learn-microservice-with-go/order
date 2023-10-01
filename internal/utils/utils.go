package utils

import (
	"math/rand"
	"time"
)

func GenerateVerificationCode(length int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	code := make([]byte, length)
	for i := 0; i < length; i++ {
		code[i] = byte(rand.Intn(10) + 48)
	}
	return string(code)
}
