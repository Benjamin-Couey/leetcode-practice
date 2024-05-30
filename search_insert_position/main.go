package search_insert_position

/* Assumes that:
1 <= nums.length <= 10^4
-10^4 <= nums[i] <= 10^4
nums contains distinct values sorted in ascending order.
-10^4 <= target <= 10^4 */
func SearchInsert(nums []int, target int) int {

	if len(nums) < 1 {
		return 0
	}

	lower_bound := 0
	upper_bound := len(nums)

	index := (upper_bound - lower_bound) / 2 + lower_bound
	val := nums[ index ]

	for upper_bound - lower_bound > 1 {
		if val == target {
			return index
		} else if val < target {
			lower_bound = index
		} else {
			upper_bound = index
		}

		index = (upper_bound - lower_bound) / 2 + lower_bound
		val = nums[ index ]
	}

	if val < target {
		return index + 1
	} else {
		return index
	}
}
