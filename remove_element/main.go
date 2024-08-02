/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/remove-element/description/
*/
package remove_element

/*
RemoveElement modifies nums to remove all occurrences of val and returns the
number of remaining elements.
RemoveElement assumes that:
0 <= nums.length <= 100,
0 <= nums[i] <= 50,
and 0 <= val <= 100.
*/
func RemoveElement(nums []int, val int) int {
	swap_index := len(nums) - 1
	removed_elements := 0
	for index := 0; index < swap_index; index++ {
		for nums[index] == val && index <= swap_index {
			nums[index] = nums[swap_index]
			nums[swap_index] = val
			swap_index -= 1
			removed_elements += 1
		}
	}
	return len(nums) - removed_elements
}
