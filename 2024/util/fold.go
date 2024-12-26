package util

func FoldL[T, U any](f func(U, T) U, init U, xs []T) U {
	acc := init
	for _, x := range xs {
		acc = f(acc, x)
	}
	return acc
}
