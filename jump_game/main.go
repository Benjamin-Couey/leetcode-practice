/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/jump-game/description/
*/
package jump_game

/*
CanJump reports whether you can jump to the last index of nums. nums[i]
represents your maximum jump length at that position. You start at nums[0].
CanJump assumes that:
1 <= nums.length <= 10^4,
and 0 <= nums[i] <= 10^5.
*/
func CanJump(nums []int) bool {
	/*
	You only can't reach the last index if there is is a zero which can't be
	jumped over. So, look for a zero that is further than any previous index
	lets you jump and return false in that case.
	*/
	max_furthest_safe_zero := -1
	for index := 0; index < len(nums)-1; index++ {
		num := nums[index]
		if num == 0 {
			if max_furthest_safe_zero <= index {
				return false
			}
		} else {
			if index+num >= len(nums)-1 {
				return true
			}
			furthest_safe_zero := index + num - 1
			if furthest_safe_zero > max_furthest_safe_zero {
				max_furthest_safe_zero = furthest_safe_zero
			}
		}
	}
	return true
}
