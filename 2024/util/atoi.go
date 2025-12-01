package util

import (
	"strconv"
)

type stringish interface {
	~string | ~rune
}

func MustAtoi[X stringish](x X) int {
	i, err := strconv.Atoi(string(x))
	if err != nil {
		panic(err)
	}

	return i
}
func MustAtoi64[X stringish](x X) int64 {
	i, err := strconv.ParseInt(string(x), 10, 64)
	if err != nil {
		panic(err)
	}

	return i
}
