package util

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	random *rand.Rand
)

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandInt generate a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + random.Int63n(max-min+1)
}

// RandomOwner generate a random name
func RandomOwner() string {
	owners := []string{"atanda", "ismail", "sultan", "munirah", "omowunmi", "victor", "nasu", "mujeeb", "samod", "rasheed", "osho", "agenla", "adeleke"}
	n := len(owners)
	return owners[random.Intn(n)]
}

// RandomMoney generate a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 100)
}

func RandomEmail() string {
	return fmt.Sprintf("%s%d@gmail.com", RandomOwner(), random.Intn(1000))
}
