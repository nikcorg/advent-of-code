package slices

import (
	"reflect"
	"testing"
)

func TestBinsert(t *testing.T) {
	cases := []struct {
		Base     []int
		Insert   int
		Expected []int
	}{
		{[]int{}, 4, []int{4}},
		{[]int{1, 3, 5}, 4, []int{1, 3, 4, 5}},
		{[]int{1, 3, 4}, 2, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 5}, 4, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 5}, 5, []int{1, 2, 3, 5}},
		{[]int{1, 2, 3, 5}, 2, []int{1, 2, 3, 5}},
	}

	for _, c := range cases {
		outcome := binsert(c.Base, c.Insert)

		if !reflect.DeepEqual(c.Expected, outcome) {
			t.Errorf("expected %d, got: %d", c.Expected, outcome)
		}
	}
}
