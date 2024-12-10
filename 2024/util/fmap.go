package util

func Fmap[T, U any](f func(T) U, xs []T) []U {
	us := make([]U, len(xs))

	for i := range xs {
		us[i] = f(xs[i])
	}

	return us
}
