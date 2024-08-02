/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/sqrtx/description/
*/
package sqrt

/*
MySqrt returns the square root of x, calculated using
https://en.wikipedia.org/wiki/Newton's_method
MySqrt assumes that 0 <= x <= 2^31 - 1.
*/
func MySqrt(x int) int {
	guess := float32(1)
	for x != int(guess*guess) {
		guess = (guess + float32(x)/guess) / float32(2)
	}
	return int(guess)
}
