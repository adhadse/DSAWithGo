package mathematics

// GCD calculates Greatest Common Divisor
// iteratively using remainder.
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return absInt(a)
}
