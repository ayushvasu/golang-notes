package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabate = "abcdefghijklmnopqrstuvwxyz"

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	rand.New(source)
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabate)
	for i := 0; i < n; i++ {
		c := alphabate[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 100000)
}

func RandomCurrency() string {
	currenices := []string{EUR, USD, INR, IDR}
	n := len(currenices)
	return currenices[rand.Intn(n)]
}
