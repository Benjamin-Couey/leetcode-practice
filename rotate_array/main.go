package rotate_array

/* Assumes that:
1 <= nums.length <= 10^5
-23^1 <= nums[i] <= 23^1 - 1
0 <= k <= 10^5 */
func Rotate(nums []int, k int) {
	rotate_index := len(nums) - k
	copy( nums, append( nums[rotate_index:], nums[:rotate_index]... ) )
}

/* Second implementation that rotates array in place with O(1) extra space. */
func AltRotate(nums []int, k int) {
	iteration := 0
	left_index := iteration * k
	right_index := len(nums) - k
	for left_index < right_index {
		for right_index < len(nums) {
			temp := nums[ left_index ]
			nums[ left_index ] = nums[ right_index ]
			nums[ right_index ] = temp
			left_index++
			right_index++
		}

		iteration++
		left_index = iteration * k
		right_index = len(nums) - k
	}
}
