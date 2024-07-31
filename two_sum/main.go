package two_sum

/* Assumes that:
2 <= nums.length <= 10^4
-10^9 <= nums[i] <= 10^9
-10^9 <= target <= 10^9

Each input HAS exactly one solution, and the same element twice cannot be used
twice. */
func TwoSum(nums []int, target int) []int {

	pair_to_index := make(map[int]int)

	for index, num := range nums {
		pair_index, exists := pair_to_index[num]
		if exists {
			return []int{pair_index, index}
		}
		pair_to_index[target-num] = index
	}

	return nil
}
