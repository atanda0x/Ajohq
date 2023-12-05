package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

// RandomInit func generate a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString fun generate a random string of len n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner func generate a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney func generate a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency func generate a random curency over six continent and ninteen countries
func RandomCurrency() string {
	currencies := []string{XOF, XAF, ZAR, EUR, CNY, JPY, INR, GBP, CHF, USD, CAD, MXN, AUD, NZD, BRL, ARS, CLP, NGN}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@gmil.com", RandomString(6))
}
