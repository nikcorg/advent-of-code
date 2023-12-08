package util

// GCD and LCM from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, rest ...int) int {
	lcm := a * b / GCD(a, b)

	for _, n := range rest {
		lcm = LCM(lcm, n)
	}

	return lcm
}
