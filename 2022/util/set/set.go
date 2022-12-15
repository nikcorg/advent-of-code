package set

type Set[T comparable] map[T]struct{}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) Has(x T) bool {
	_, ok := s[x]
	return ok
}

func (s Set[T]) Add(x T) {
	s[x] = struct{}{}
}

func (s Set[T]) Remove(x T) {
	delete(s, x)
}

func (s Set[T]) Intersection(x Set[T]) Set[T] {
	a, b := s, x
	if a.Size() < b.Size() {
		a, b = b, a
	}

	i := New[T]()
	for ka := range a {
		if b.Has(ka) {
			i.Add(ka)
		}
	}

	return i
}

func New[T comparable](inp ...T) Set[T] {
	m := make(Set[T])

	for _, c := range inp {
		m[c] = struct{}{}
	}

	return m
}
