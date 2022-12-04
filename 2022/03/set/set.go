package set

type Set[T comparable] map[T]struct{}

func (s Set[T]) Size() int {
	return len(s)
}

func New[T comparable](inp []T) Set[T] {
	m := make(Set[T])

	for _, c := range inp {
		m[c] = struct{}{}
	}

	return m
}
