package utils

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomGenre() string {
	names := []string{"horror", "comedy", "drama", "action", "romance", "sci-fi", "fantasy"}
	name := names[rand.Intn(len(names))]

	return name
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		sb.WriteByte(alphabet[rand.Intn(k)])
	}
	return sb.String()
}

func RandomBook() string {
	return RandomString(10)
}

func RandomName() string {
	return RandomString(6)
}
