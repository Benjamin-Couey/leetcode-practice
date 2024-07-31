package single_number

/* Assumes that:
1 <= nums.length <= 3 * 10^4
-3 * 10^4 <= nums[i] <= 3 * 10^4
Each element in the array appears twice except for one element which appears
only once. */
func SingleNumber(nums []int) int {
	var mask int

	for _, num := range nums {
		mask = mask ^ num
	}

	return mask
}
