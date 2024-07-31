package jump_game

/* Assumes that:
1 <= nums.length <= 10^4
0 <= nums[i] <= 10^5 */
func CanJump(nums []int) bool {
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
