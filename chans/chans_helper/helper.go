package chans_helper

import (
	"math/rand"
	"time"
)

func RandomNumberGenerator(pool int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(pool)
}
