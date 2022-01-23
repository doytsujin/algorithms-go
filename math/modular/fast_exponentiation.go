package modular

// Exponentiation returns base^exponent
func FastExponentiation(base, exponent int64) int64 {
	if exponent == 0 {
		return 1
	} else if exponent == 1 {
		return base
	}

	r := FastExponentiation(base, exponent/2)

	if exponent%2 == 0 {
		return r * r
	} else {
		return r * base * r
	}
}
