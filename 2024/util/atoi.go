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
