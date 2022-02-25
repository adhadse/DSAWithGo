package bit_magic

func CountSetbitBrianKerningham(a int) int {
	count := 0
	for a != 0 {
		a &= a - 1
		count++
	}
	return count
}
