package util

import (
	"crypto/rand"
)

func GenerateRandomString(length int) (string, error) {
	possible := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	v := make([]byte, length)
	_, err := rand.Read(v)
	if err != nil {
		return "", err
	}

	result := make([]byte, length)
	for i, b := range v {
		result[i] = possible[int(b)%len(possible)]
	}

	return string(result), nil
}
