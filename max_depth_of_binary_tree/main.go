/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
*/
package max_depth_of_binary_tree

import (
	"leetcode/utils"
)

/*
MaxDepth returns the maximum depth of the binary tree starting at root.
MaxDepth assumes that:
the number of nodes in the tree is in the range [0, 10^4],
-100 <= Node.val <= 100,
and Node.val != 0.
*/
func MaxDepth(root *utils.TreeNode) int {

	if root == nil {
		return 0
	}

	left_depth := MaxDepth(root.Left)
	right_depth := MaxDepth(root.Right)

	return 1 + max(left_depth, right_depth)
}
