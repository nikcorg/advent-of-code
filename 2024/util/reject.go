package util

func Reject[T any](f func(T) bool, xs []T) []T {
	ys := []T{}

	for _, x := range xs {
		if !f(x) {
			ys = append(ys, x)
		}
	}

	return ys
}
