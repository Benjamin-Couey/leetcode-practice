/* Package includes implementation and tests for problem described at
https://leetcode.com/problems/jump-game-ii/description/ */
package jump_game_ii

/* Jump returns the minimum number of jumps needed to reach the last index to
reach the last index of nums. nums[i] represents you maximum jump length at that
position. You start at nums[0].
Jump assumes that:
1 <= nums.length <= 104,
0 <= nums[i] <= 1000,
and it's guaranteed that you can reach nums[n - 1]. */
func Jump(nums []int) int {

	jumps := 0
	index := 0

	for index < (len(nums) - 1) {
		// Always jump to the end if possible
		if index+nums[index] >= (len(nums) - 1) {
			return jumps + 1
		}
		// Find next index that gets us the furthest
		max_jump_dist := nums[index]
		best_index := index + 1
		best_jump := 1
		for jump_dist := 1; jump_dist <= max_jump_dist; jump_dist++ {
			jump_index := index + jump_dist
			total_dist := jump_dist + nums[jump_index]
			if total_dist > best_jump {
				best_jump = total_dist
				best_index = jump_index
			}
		}
		jumps++
		index = best_index
	}

	return jumps
}
