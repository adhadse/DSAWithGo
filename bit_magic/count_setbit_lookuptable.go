package bit_magic

func CountSetBitLookUpTable(a int) int {
	bitsettable := make([]int, 256)
	for i := 0; i < 256; i++ {
		bitsettable[i] = (i & 1) + bitsettable[i/2]
	}

	return bitsettable[a&0xFF] +
		bitsettable[(a>>8)&0xFF] +
		bitsettable[(a>>16)&0xFF] +
		bitsettable[(a>>24&0xFF)]
}
