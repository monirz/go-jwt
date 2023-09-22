package utils

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SplitIntoChunks(data string, maxTokens int) []string {
	var chunks []string
	runes := []rune(data)

	for len(runes) > 0 {
		if len(runes) <= maxTokens {
			chunks = append(chunks, string(runes))
			break
		}
		chunks = append(chunks, string(runes[:maxTokens]))
		runes = runes[maxTokens:]
	}
	return chunks
}

// func GetRandomNumber() int {

// 	rand.Seed(time.Now().UnixNano())
// 	min := 10000000
// 	max := 99999999
// 	return rand.Intn(max-min+1) + min
// }

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func RandStr(length int) string {

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
