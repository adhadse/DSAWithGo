package bit_magic

// CountSetBit counts the number of bits are set (1) in a
// given number
func CountSetBit(a int) int {
	if a == 0 {
		return 0
	} else {
		return (a & 1) + CountSetBit(a>>1)
	}
}
