package util

func CloneSlice[X any](xs []X) []X {
	xs0 := make([]X, len(xs))
	copy(xs0, xs)
	return xs0
}
