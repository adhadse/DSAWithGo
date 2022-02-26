package recursion

// TailRecursive is a special form of recursion where
// the last operation of a function is a recursive call.
func TailRecursive(n, a int) int {
	if n == 0 {
		return a
	}
	return TailRecursive(n-1, n*a)
}

// NonTailRecursive is not tail recursive
// because the value returned by NonTailRecursive(n-1)
// is used in NonTailRecursive(n) and call to NonTailRecursive(n-1)
// is not the last thing done by NonTailRecursive(n)
func NonTailRecursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * NonTailRecursive(n-1)
}
