package slices

type SortedIntSlice []int

// Binsert inserts the value `v` into the IntSlice returning a new sorted IntSlice
func (s SortedIntSlice) Insert(v int) SortedIntSlice {
	return binsert(s, v)
}

func (s SortedIntSlice) First() int {
	return s[0]
}

func (s SortedIntSlice) Last() int {
	return s[len(s)-1]
}

func (s SortedIntSlice) Append(v int) []int {
	return append(s, v)
}

func (s SortedIntSlice) Max() int {
	return s.Last()
}

func (s SortedIntSlice) Min() int {
	return s.First()
}

func (s SortedIntSlice) Unshift() (int, SortedIntSlice) {
	return s.First(), s[1:]
}

func (s SortedIntSlice) Len() int {
	return len(s)
}
