package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackPush(t *testing.T) {
	s := New[int]()

	s.Push(3, 2, 1)

	assert.Equal(t, 3, s.Size())
}

func TestStackPop(t *testing.T) {
	s := New[int]()

	s.Push(3, 2, 1)

	x := s.Pop()

	assert.Equal(t, 3, x)
}

func TestStackPopN(t *testing.T) {
	s := New[int]()

	s.Push(3, 2, 1)

	assert.Equal(t, []int{3, 2, 1}, s.PopN(3))

	s = New[int]()
	s.Push(3)
	s.Push(2)
	s.Push(1)

	assert.Equal(t, []int{1, 2, 3}, s.PopN(3))
}

func TestStackPeek(t *testing.T) {
	s := New[int]()

	s.Push(3, 2, 1)
	assert.Equal(t, 3, s.Peek())
}
