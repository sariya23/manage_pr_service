//go:build integrations

package random

import "math/rand"

// RandInt случайное число в диапазоне [min, max]
func RandInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}
