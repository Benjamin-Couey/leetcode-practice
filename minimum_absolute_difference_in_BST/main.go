/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/minimum-absolute-difference-in-bst/description/
*/
package minimum_absolute_difference_in_BST

import (
	"leetcode/utils"
)

// The maximum difference between values in the tree.
const MAX_VAL int = 100000

/*
GetMinimumDifference returns the minimum absolute difference between any two nodes
in the tree starting at root.
GetMinimumDifference assumes that:
the number of nodes in the tree is in the range [2, 10^4],
and 1 <= TreeNode.val <= 10^5.
*/
func GetMinimumDifference(root *utils.TreeNode) int {
	return recursiveGetMinimumDifference(root, root)
}

/*
recursiveGetMinimumDifference traverses the tree and passes each node to
recursiveCompareToTree, then returns the minimum of the returned differences.
*/
func recursiveGetMinimumDifference(root *utils.TreeNode, cursor *utils.TreeNode) int {
	if cursor == nil {
		return MAX_VAL
	}
	diff := recursiveCompareToTree(root, cursor)
	left_diff := recursiveGetMinimumDifference(root, cursor.Left)
	if left_diff < diff {
		diff = left_diff
	}
	right_diff := recursiveGetMinimumDifference(root, cursor.Right)
	if right_diff < diff {
		diff = right_diff
	}
	return diff
}

/*
recursiveCompareToTree returns the minimum distance between a given node
and all other nodes in the tree.
*/
func recursiveCompareToTree(cursor *utils.TreeNode, node *utils.TreeNode) int {
	diff := MAX_VAL
	if cursor == nil {
		return diff
	}
	if cursor != node {
		diff = absDiff(cursor.Val, node.Val)
	}
	left_diff := recursiveCompareToTree(cursor.Left, node)
	if left_diff < diff {
		diff = left_diff
	}
	right_diff := recursiveCompareToTree(cursor.Right, node)
	if right_diff < diff {
		diff = right_diff
	}
	return diff
}

func absDiff(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
