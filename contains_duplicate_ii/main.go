package contains_duplicate_ii

/* Assumes that:

1 <= nums.length <= 10^5
-10^9 <= nums[i] <= 10^9
0 <= k <= 10^5 */
func ContainsNearbyDuplicate(nums []int, k int) bool {

	value_to_last_index := make(map[int]int)

	for index, value := range nums {
		last_index, exists := value_to_last_index[value]
		if exists && index-last_index <= k {
			return true
		}
		value_to_last_index[value] = index
	}

	return false
}
