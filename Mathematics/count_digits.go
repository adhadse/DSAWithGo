package Mathematics

import (
	"strconv"
)
func LenOfString(n int) int {
	// Return length of string
	// convert integer to string and get the length of the string
	return len(strconv.Itoa(absInt(n)))
}
