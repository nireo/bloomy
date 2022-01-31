package bloomy_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/nireo/bloomy"
)

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func TestBasic(t *testing.T) {
	str := []byte("hello world")
	bf := bloomy.New(10000, 0.01)

	bf.Insert(str)

	if !bf.Contains(str) {
		t.Error("couldn't find string in bloom filter")
	}
}

func TestLargeInput(t *testing.T) {
	bf := bloomy.New(10000, 0.001)
	rand.Seed(time.Now().Unix())
	iterCount := 10000
	for i := 0; i < iterCount; i += 1 {
		randString := randomString(16)
		bf.Insert([]byte(randString))

		if !bf.Contains([]byte(randString)) {
			t.Error("couldn't find string in bloom filter")
		}
	}
}
