package util

import "strconv"

func ParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
