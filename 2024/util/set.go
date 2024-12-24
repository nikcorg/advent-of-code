package util

import "maps"

type void struct{}

type Set[T comparable] map[T]void

func NewSet[T comparable](init ...T) Set[T] {
	s := make(Set[T])

	for _, v := range init {
		s.Add(v)
	}

	return s
}

func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Values() []T {
	xs := make([]T, len(s))
	i := 0
	for v := range maps.Keys(s) {
		xs[i] = v
		i++
	}
	return xs
}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}
