/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/description/
*/
package remove_duplicates_from_sorted_array_ii

/*
RemoveDuplicates remove duplicate elements from nums such that each element appears
no more than twice and returns the remaining number of elements in nums. Put
another way, RemoveDuplicates returns k and modifies nums such that the first k
elements of nums are elements appearing no more than twice in the order they were
in nums initially.
RemoveDuplicates assumes that:
1 <= nums.length <= 3 * 10^4,
-10^4 <= nums[i] <= 10^4,
and nums is sorted in non-decreasing order.
*/
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
