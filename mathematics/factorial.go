package mathematics

func Factorial(x int) int {
	// find factorial by recursion.
	if x == 0 || x == 1 {
		return 1
	} else {
		return x * Factorial(x-1)
	}
}
