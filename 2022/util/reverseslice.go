package util

// ReverseSlice reverses a slice in-place
func ReverseSlice[T any](xs []T) []T {
	for l, r := 0, len(xs)-1; l < r; l, r = l+1, r-1 {
		xs[l], xs[r] = xs[r], xs[l]
	}
	return xs
}
