package mathematics

import "math"

// LCM Lowest common factor returns
// By dividing multiplication of the two numbers by their
// Greatest common Divisor
func LCM(a, b int) int {
	return absInt(int(
		math.Floor(float64(a*b) / float64(GCD(a, b)))))
}
