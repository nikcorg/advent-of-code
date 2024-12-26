package util

func Last[T any](xs []T) T {
	return xs[len(xs)-1]
}
