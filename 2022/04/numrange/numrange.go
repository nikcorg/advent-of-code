package intrange

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

var (
	errInvalidInput = errors.New("invalid input for int range")
)

type rangeable interface {
	constraints.Integer | constraints.Float
}

type NumRange[T rangeable] struct {
	l, u T
}

func (r NumRange[T]) Contains(x NumRange[T]) bool {
	if r.l <= x.l && x.u <= r.u {
		return true
	}
	return false
}

func (r NumRange[T]) Overlaps(x NumRange[T]) bool {
	if x.l > r.u || r.l > x.u {
		return false
	}

	return true
}

func (r NumRange[T]) String() string {
	return fmt.Sprintf("%v-%v", r.l, r.u)
}

func (r NumRange[T]) Lower() T {
	return r.l
}

func (r NumRange[T]) Upper() T {
	return r.u
}

func New[T rangeable](l, u T) NumRange[T] {
	if l > u {
		l, u = u, l
	}
	return NumRange[T]{l, u}
}

func FromString[T rangeable](s string, f func(string) (T, error)) (NumRange[T], error) {
	splits := strings.Split(s, "-")
	if len(splits) != 2 {
		return NumRange[T]{}, errInvalidInput
	}

	l, err := f(splits[0])
	if err != nil {
		return NumRange[T]{}, err
	}

	u, err := f(splits[1])
	if err != nil {
		return NumRange[T]{}, err
	}

	return New(l, u), nil
}
