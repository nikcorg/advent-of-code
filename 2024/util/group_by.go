package util

func GroupBy[A any, C comparable](f func(A) C, as []A) map[C][]A {
	gs := map[C][]A{}

	for _, a := range as {
		c := f(a)

		if g, ok := gs[c]; ok {
			gs[c] = append(g, a)
		} else {
			gs[c] = []A{a}
		}
	}

	return gs
}
