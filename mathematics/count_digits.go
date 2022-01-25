package mathematics

import (
	"strconv"
)

// LenOfString return length of string
// convert integer to string and get the length of the string
func LenOfString(n int) int {
	return len(strconv.Itoa(absInt(n)))
}
