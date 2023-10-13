package utils

import "math/rand"

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[RandomInt(len(letters))]
	}
	return string(b)
}

func RandomInt(n int) int {
	return rand.Intn(n)
}
