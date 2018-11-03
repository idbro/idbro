package middleware

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandomString(t *testing.T) {
	randomString := getRandomString()
	log.Println(randomString)
	assert.Equal(t, len(randomString), strLen, "random string length wrong")
}

func BenchmarkGetRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = getRandomString()
	}
}
