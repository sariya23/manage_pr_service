package helpers

import "slices"

func Filter[T comparable](a []T, b T) []T {
	res := make([]T, 0, len(a)-1)
	for _, x := range a {
		if x != b {
			res = append(res, x)
		}
	}
	return res
}

func Filters[T comparable](set []T, subset []T) []T {
	if len(set) < len(subset) {
		panic("Set less than subset")
	}
	res := make([]T, 0, len(set)-len(subset))
	for _, x := range set {
		if !slices.Contains(subset, x) {
			res = append(res, x)
		}
	}
	return res
}
