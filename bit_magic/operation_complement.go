package bit_magic

// Complement finds the One's complement of a number
// using the same ^ operator used for XOR operation
func Complement(a int) int {
	return ^a
}
