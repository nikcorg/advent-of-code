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
		s := New(tc.in)
		assert.Equal(t, tc.size, s.Size())
	}
}
