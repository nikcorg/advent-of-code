package util

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](x, y T) T {
	if x > y {
		return y
	}
	return x
}

func Max[T constraints.Ordered](x, y T) T {
	if x < y {
		return y
	}
	return x
}
