package random

import (
	"math/rand"
	"time"
)

func Sample[T any](slice []T, n int) []T {
	if n <= 0 || len(slice) == 0 {
		return nil
	}

	if n >= len(slice) {
		result := make([]T, len(slice))
		copy(result, slice)
		return result
	}

	rand.Seed(time.Now().UnixNano())
	indices := rand.Perm(len(slice))[:n]

	result := make([]T, n)
	for i, idx := range indices {
		result[i] = slice[idx]
	}

	return result
}
