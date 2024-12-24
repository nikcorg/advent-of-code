package util

func FlatMap[T, U any](f func(T) []U, xs []T) []U {
	us := []U{}

	for _, x := range xs {
		us = append(us, f(x)...)
	}

	return us
}
