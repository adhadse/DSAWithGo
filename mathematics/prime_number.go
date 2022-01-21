package mathematics

import "math"

func IsPrime(x int) bool {
	// Return a boolean value to check if
	// there exists any factors greater than 1.
	// a number divisible only by itself and 1.
	if x < 2 {
		return false
	} else {
		for i := 2; i < int(math.Sqrt(float64(x))+1); i++ {
			if x%i == 0 {
				return false
			}
		}
	}
	return true
}
