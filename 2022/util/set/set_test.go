package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	cases := []struct {
		in   []byte
		size int
	}{
		{[]byte("hello world"), 8},
		{[]byte("aaaaaaaaaa"), 1},
		{[]byte("aaaaa aaaaa"), 2},
	}

	for _, tc := range cases {
		s := New(tc.in...)
		assert.Equal(t, tc.size, s.Size())
	}
}

func TestSetIntersection(t *testing.T) {
	cases := []struct {
		a, b []string
		i    Set[string]
	}{
		{[]string{"a", "b", "c"}, []string{"b", "c", "d"}, New("b", "c")},
		{[]string{"a", "b", "c"}, []string{"d", "e", "f"}, New[string]()},
	}

	for _, tc := range cases {
		a, b := New(tc.a...), New(tc.b...)
		ia := a.Intersection(b)
		ib := b.Intersection(a)

		assert.Equal(t, tc.i, ib)
		assert.Equal(t, ia, ib)
	}
}
