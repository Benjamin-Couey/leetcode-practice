package is_height_balanced

import (
	"leetcode/max_depth_of_binary_tree"
	"leetcode/utils"
)

func IsHeightBalanced(root *utils.TreeNode) bool {

	if root == nil {
		return true
	}

	left_depth := max_depth_of_binary_tree.MaxDepth(root.Left)
	right_depth := max_depth_of_binary_tree.MaxDepth(root.Right)

	diff := left_depth - right_depth
	if right_depth > left_depth {
		diff = right_depth - left_depth
	}

	if diff > 1 {
		return false
	}

	return IsHeightBalanced(root.Left) && IsHeightBalanced(root.Right)
}
