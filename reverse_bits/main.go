/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/reverse-bits/description/
*/
package reverse_bits

/*
ReverseBits reverses the bits of num and returns the result as an unsigned
32 bit integer.
ReverseBits assumes the input is a binary string of length 32.
*/
func ReverseBits(num uint32) uint32 {
	var result uint32
	for index := 0; index < 32; index++ {
		result <<= 1
		result = result + num&1
		num >>= 1
	}
	return result
}
