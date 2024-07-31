package sqrt

/* Assumes that:
0 <= x <= 2^31 - 1 */
func MySqrt(x int) int {
	guess := float32(1)
	for x != int(guess*guess) {
		guess = (guess + float32(x)/guess) / float32(2)
	}
	return int(guess)
}
