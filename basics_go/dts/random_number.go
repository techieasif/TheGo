package dts

import (
	"math"
	"math/rand"
)

func generateRandomNumber() int {
	rand.Seed(10)
	return rand.Intn(11)
}

func showValueOfPi() float64 {
	return math.Pi
}
