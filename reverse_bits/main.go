package reverse_bits

/* Assumes the input is a binary string of length 32. */
func ReverseBits(num uint32) uint32 {
	var result uint32
	for index := 0; index < 32; index++ {
		result <<= 1
		result = result + num&1
		num >>= 1
	}
	return result
}
