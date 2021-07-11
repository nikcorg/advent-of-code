package utils

// source: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b uint64) uint64 {
	for b != 0 {
		b, a = a%b, b
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b uint64, integers ...uint64) uint64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
