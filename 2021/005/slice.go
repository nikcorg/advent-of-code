package main

type filterfn[T any] func(x T) bool

func filter[T any](xs []T, f filterfn[T]) []T {
	out := []T{}
	for _, x := range xs {
		if f(x) {
			out = append(out, x)
		}
	}
	return out
}

type comparefn[T any, U comparable] func(x T) U

func uniq[T any, U comparable](xs []T, f comparefn[T, U]) []T {
	uxs := make(map[U]struct{})
	out := []T{}

	for _, x := range xs {
		cmp := f(x)
		if _, ok := uxs[cmp]; !ok {
			uxs[cmp] = struct{}{}
			out = append(out, x)
		}
	}

	return out
}
