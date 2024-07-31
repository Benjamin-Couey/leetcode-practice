package remove_element

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
