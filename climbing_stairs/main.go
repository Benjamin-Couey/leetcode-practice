/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/climbing-stairs/description/
*/
package climbing_stairs

/*
ClimbStairs retrusn how many distinct way a staircase with n steps can be
climbed, 1 or 2 steps at a time.
ClimbStairs assumes that 1 <= n <= 45.
*/
func ClimbStairs(n int) int {
	sum := 0
	n_minus_one := 1
	n_minus_two := 0

	for index := 0; index < n; index++ {
		sum = n_minus_one + n_minus_two
		n_minus_two = n_minus_one
		n_minus_one = sum
	}

	return sum
}
