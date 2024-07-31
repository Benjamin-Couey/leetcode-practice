package remove_duplicates_from_sorted_array_ii

/*Assumes that:
1 <= nums.length <= 3 * 10^4
-10^4 <= nums[i] <= 10^4
nums is sorted in non-decreasing order. */
func RemoveDuplicates(nums []int) int {
	count := 1
	current_num := nums[0]
	index := 1
	upper_bound := len(nums)
	for index < upper_bound {
		if nums[index] == current_num {
			count++
			if count > 2 {
				nums = append(nums[:index], nums[index+1:]...)
				upper_bound--
			} else {
				index++
			}
		} else {
			count = 1
			current_num = nums[index]
			index++
		}
	}
	return index
}
