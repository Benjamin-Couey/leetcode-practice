/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/number-of-1-bits/description/
*/
package number_of_one_bits

/*
HammingWeight returns the number of 1 bits in the binary representation of n
(the Hamming weight).
HammingWeight assumes that 1 <= n <= 2^31 - 1.
*/
func HammingWeight(n int) int {
	count := 0
	for index := 0; index < 32; index++ {
		if n&1 == 1 {
			count++
		}
		n >>= 1
	}
	return count
}
