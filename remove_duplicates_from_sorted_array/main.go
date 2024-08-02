/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
*/
package remove_duplicates_from_sorted_array

/*
RemoveDuplicates remove duplicate elements from nums and returns the remaining
number of elements in nums. Put another way, RemoveDuplicates returns k and
modifies nums such that the first k elements of nums are unique elements in the
order they were in nums initially.
RemoveDuplicates assumes that:
1 <= nums.length <= 3 * 10^4,
-100 <= nums[i] <= 100,
and nums is sorted in non-decreasing order.
*/
func RemoveDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	insert_index := 1
	last_val := nums[0]

	for index := 1; index < len(nums); index++ {
		if nums[index] != last_val {
			nums[insert_index] = nums[index]
			insert_index += 1
			last_val = nums[index]
		}
	}

	return insert_index
}
