package util

import "strconv"

func MustAtoi(x string) int {
	i, err := strconv.Atoi(x)
	if err != nil {
		panic(err)
	}

	return i
}
