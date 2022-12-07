package stack

type Stack[T any] struct {
	stack []T
}

func New[T any]() Stack[T] {
	return Stack[T]{}
}

func (s *Stack[T]) Size() int {
	return len(s.stack)
}

func (s *Stack[T]) Push(vs ...T) {
	s.stack = append(vs, s.stack...)
}

func (s *Stack[T]) Pop() T {
	var v T
	v, s.stack = s.stack[0], s.stack[1:]
	return v
}

func (s *Stack[T]) PopN(n int) []T {
	var vs []T
	vs, s.stack = s.stack[0:n], s.stack[n:]
	return vs
}

func (s *Stack[T]) Peek() T {
	var x T
	if len(s.stack) > 0 {
		x = s.stack[0]
	}
	return x
}
