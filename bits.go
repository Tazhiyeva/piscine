package piscine

func reverseBits(oct byte) byte {
	var result byte

	for i := 0; i < 8; i++ {
		result = (result << 1) | (oct & 1)
		oct >>= 1
	}
	return result
}

func swapBits(oct byte) byte {
	return (oct<<4 | oct>>4)

}
