//go:build integrations

package random

import (
	"math/rand"
	"time"
)

func Choice[T any](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}

	rand.Seed(time.Now().UnixNano())
	return slice[rand.Intn(len(slice))]
}
