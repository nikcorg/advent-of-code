package util

func Fmap[T any, U any](f func(T) U, xs []T) []U {
	ys := make([]U, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}
