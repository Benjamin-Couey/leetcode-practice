package climbing_stairs

/* Assumes that:
1 <= n <= 45 */
func ClimbStairs(n int) int {
	sum := 0
	n_minus_one := 1
	n_minus_two := 0

	for index := 0; index < n; index++ {
		sum = n_minus_one +	n_minus_two
		n_minus_two = n_minus_one
		n_minus_one = sum
	}

	return sum
}
