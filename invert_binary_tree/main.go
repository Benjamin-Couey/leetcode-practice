/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/invert-binary-tree/description/
*/
package invert_binary_tree

import (
	"leetcode/utils"
)

/*
InvertTree inverts binary tree starting at root and returns root.
InvertTree assumes that:
the number of nodes in the tree is in the range [0, 100],
-100 <= Node.val <= 100 and Node.val != 0.
*/
func InvertTree(root *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}

	left_child := root.Left
	root.Left = InvertTree(root.Right)
	root.Right = InvertTree(left_child)
	return root
}
