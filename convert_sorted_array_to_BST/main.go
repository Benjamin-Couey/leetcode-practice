package convert_sorted_array_to_BST

import (
	"leetcode/utils"
)

func SortedArrayToBST(nums []int) *utils.TreeNode {
	if len(nums) < 1 {
		return nil
	}

	index := len(nums) / 2

	root := utils.TreeNode{nums[index], nil, nil}
	root.Left = SortedArrayToBST(nums[:index])
	root.Right = SortedArrayToBST(nums[index+1:])

	return &root
}
