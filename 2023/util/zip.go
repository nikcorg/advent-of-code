package util

func Zip[T any](xs, ys []T) [][]T {
	l := min(len(xs), len(ys))
	zs := make([][]T, l)

	for l > 0 {
		zs[l-1] = []T{xs[l-1], ys[l-1]}
		l--
	}

	return zs
}
