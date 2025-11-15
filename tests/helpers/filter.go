package helpers

func Filter[T comparable](a []T, b T) []T {
	res := make([]T, 0, len(a)-1)
	for _, x := range a {
		if x != b {
			res = append(res, x)
		}
	}
	return res
}

//func Filters[T comparable](a []T, b []T) []T {
//	res := make([]T, 0, len(a) - 1)
//	for _, x := range a {
//		if slices.Contains() {
//			res = append(res, x)
//		}
//	}
//	return res
//}
//
