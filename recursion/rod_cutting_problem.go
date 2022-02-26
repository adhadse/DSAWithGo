package recursion

import "math"

// RodCutting is a problem where a given rod of length n
// and price list of all piece length < n; we need to find
// the maximum value obtainable by selling up cut rod pieces.
func RodCutting(pricePerLength []int, length int) int {
	if length <= 0 {
		return 0
	}
	maxVal := -math.MaxInt64
	for i := 0; i < length; i++ {
		maybeMaxVal := pricePerLength[i] + RodCutting(pricePerLength, length-i-1)
		if maxVal <= maybeMaxVal {
			maxVal = maybeMaxVal
		}
	}
	return maxVal
}
