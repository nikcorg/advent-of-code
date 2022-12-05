package util

import "strconv"

func MustAtoi(x string) int {
	n, err := strconv.Atoi(x)
	if err != nil {
		panic(err)
	}
	return n
}
