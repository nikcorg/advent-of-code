package util

func Fmap[X, Y any](f func(X) Y, xs []X) []Y {
	ys := make([]Y, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}
