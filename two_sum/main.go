/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/two-sum/description/
*/
package two_sum

/*
TwoSum returns two indicies of two numbers in nums that sum to target. The
same element twice cannot be used twice.
TwoSum assumes that:
2 <= nums.length <= 10^4,
-10^9 <= nums[i] <= 10^9,
-10^9 <= target <= 10^9,
each input HAS exactly one solution.
*/
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
