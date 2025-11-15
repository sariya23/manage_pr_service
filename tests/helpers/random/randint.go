package random

import (
	"math/rand"
)

// RandInt случайное число в диапазоне [min, max]
func RandInt(start, end int) int {
	return rand.Intn(end-start+1) + start
}
