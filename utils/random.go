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

func randomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		sb.WriteByte(alphabet[rand.Intn(k)])
	}
	return sb.String()
}

func RandomTitle() string {
	return randomString(10)
}

func RandomPrice(min, max int) int {

	return min + rand.Intn(max-min+1)
}

func RandomBool() bool {
	bools := []bool{true, false}

	return bools[rand.Intn(len(bools))]
}

func RandomBio() string {
	return randomString(50)
}

func RandomName() string {
	return randomString(6)
}
