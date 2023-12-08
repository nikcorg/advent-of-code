package util

import "strconv"

func ParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func ParseIntWithDefault(s string, fallback int) int {
	if n, err := strconv.Atoi(s); err == nil {
		return n
	}
	return fallback
}
