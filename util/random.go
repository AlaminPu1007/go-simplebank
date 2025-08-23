package util

import (
	"math/rand"
	"strings"
)

const characters = "abcdefghijklmnopqrstuvwxyz"

// RandomInt function to get each time new random integer value
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString function to get n no of characters
func RandomString(n int) string {
	var sb strings.Builder

	k := len(characters)

	for i := 0; i < n; i++ {
		c := characters[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Get random owner name by random string generator
func RandomOwner() string {
	return RandomString(6)
}

// Get random money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// Ger random currency
func RandomCurrency() string {
	lists := []string{"BDT", "USD", "URO", "CAD"}
	n := len(lists)

	return lists[rand.Intn(n)]
}
