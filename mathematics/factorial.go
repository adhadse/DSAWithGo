package mathematics

// Factorial returns factorial of a number by recursion.
func Factorial(x int) int {
	if x == 0 || x == 1 {
		return 1
	} else {
		return x * Factorial(x-1)
	}
}
