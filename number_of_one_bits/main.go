package number_of_one_bits

/* Assumes that:
1 <= n <= 2^31 - 1 */
func HammingWeight(n int) int {
	count := 0
	for index := 0; index < 32; index++ {
		if n & 1 == 1 {
			count ++
		}
		n >>= 1
	}
	return count
}
