package util

// the `math.Mod` and `%` operator returns negative values when `a` is negative, this replacement
// returns the expected positive result
func Mod(a, b int) int {
	return (((a % b) + b) % b)
}
