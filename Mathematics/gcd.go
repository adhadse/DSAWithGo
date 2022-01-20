package Mathematics

func GCD(a, b int) int {
	// calculates Greatest Common Divisor
	// iteratively using remainder.
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
