package util

func Zip[T any](as, bs []T) [][2]T {
	cnt := Min(len(as), len(bs))
	cs := make([][2]T, cnt)
	for i := 0; i < cnt; i++ {
		cs[i] = [2]T{as[i], bs[i]}
	}
	return cs
}
