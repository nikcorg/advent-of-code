package util

func CopySlice[T any](xs []T) []T {
	ys := make([]T, len(xs))
	copy(ys, xs)
	return ys
}
