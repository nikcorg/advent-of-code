package stack

import "sync"

type Stack[T any] struct {
	stack []T
	mut   sync.RWMutex
}

func New[T any]() Stack[T] {
	return Stack[T]{}
}

func (s *Stack[T]) Size() int {
	s.mut.RLock()
	defer s.mut.RUnlock()

	return len(s.stack)
}

func (s *Stack[T]) Push(vs ...T) {
	s.mut.Lock()
	defer s.mut.Unlock()

	s.stack = append(vs, s.stack...)
}

func (s *Stack[T]) Pop() T {
	s.mut.Lock()
	defer s.mut.Unlock()

	var v T

	if len(s.stack) > 0 {
		v, s.stack = s.stack[0], s.stack[1:]
	}

	return v
}

func (s *Stack[T]) PopN(n int) []T {
	s.mut.Lock()
	defer s.mut.Unlock()

	var vs []T
	vs, s.stack = s.stack[0:n], s.stack[n:]
	return vs
}

func (s *Stack[T]) Peek() T {
	s.mut.RLock()
	defer s.mut.RUnlock()

	var x T
	if len(s.stack) > 0 {
		x = s.stack[0]
	}
	return x
}

func (s *Stack[T]) Clear() {
	s.mut.Lock()
	defer s.mut.Unlock()

	s.stack = []T{}
}

func (s *Stack[T]) Each(f func(T)) {
	s.mut.RLock()
	defer s.mut.RUnlock()

	for _, v := range s.stack {
		f(v)
	}
}
