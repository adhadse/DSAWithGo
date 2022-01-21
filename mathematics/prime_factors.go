package mathematics

import "math"

func PrimeFactors(x int) []int {
	// Return a list of prime factors of an integer by
	// first checking if the modulo of the value of d and number is equal to 0.
	var divisors []int
	var primeFactors []int
	for d := 2; d < int(math.Floor(float64(x)/2))+1; d++ {
		if x%d == 0 {
			divisors = append(divisors, d)
		}
	}
	// Then applying the for loop inside the for loop to count a factor only once.
	for _, d := range divisors {
		for _, od := range divisors {
			if od != d {
				if d%od != 0 {
					continue
				} else {
					break
				}
			}
			primeFactors = append(primeFactors, d)
		}
	}
	return primeFactors
}
