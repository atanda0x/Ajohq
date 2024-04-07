package util

import (
	"math/rand"
	"strings"
	"time"
)

var alphabet = "abcdefghijklmnopqrstuvxyz"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

// RandInt generate a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generate a random string of len n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generate a random name
func RandomOwner() string {
	return RandomString(5)
}

// RandomMoney generate a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 100)
}
