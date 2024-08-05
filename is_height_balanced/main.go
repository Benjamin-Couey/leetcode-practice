/*
Package includes utility function for checking if a tree is height balanced.
While a utility function and not the answer to a leetcode problem, it relies on
the implementatino of max_depth_of_binary_tree and so is seperate from utils
to avoid a circular import.
*/
package is_height_balanced

import (
	"leetcode/max_depth_of_binary_tree"
	"leetcode/utils"
)

/*
IsHeightBalanced reports whether the tree starting at root is height balanced.
That is, the height of the left and the right subtree of any node differ by not
more than one.
*/
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
