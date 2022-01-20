package Mathematics

import "math"

func LCM(a, b int) int {
	// Lowest common factor
	// By dividing multiplication of the two numbers by their
	// Greatest common Divisor
	return absInt(int(
		math.Floor(float64(a*b) / float64(GCD(a, b)))))
}
