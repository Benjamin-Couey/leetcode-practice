package minimum_absolute_difference_in_BST

import (
	"leetcode/utils"
)

const MAX_VAL int = 100000

/* Assumes that:
The number of nodes in the tree is in the range [2, 10^4].
1 <= Node.val <= 10^5 */
func GetMinimumDifference(root *utils.TreeNode) int {
	return recursiveGetMinimumDifference( root, root )
}

/* Traverse the tree and pass each node to recursiveCompareToTree. Return the
minimum of the differences returned.
root - Root of the tree being traversed.
cursor - Current node in the traversal. */
func recursiveGetMinimumDifference(root *utils.TreeNode, cursor *utils.TreeNode) int {
	if cursor == nil {
		return MAX_VAL
	}
	diff := recursiveCompareToTree( root, cursor )
	left_diff := recursiveGetMinimumDifference( root, cursor.Left )
	if left_diff < diff {
		diff = left_diff
	}
	right_diff := recursiveGetMinimumDifference( root, cursor.Right )
	if right_diff < diff {
		diff = right_diff
	}
	return diff
}

/* Given the root of a tree and a node, calculate the minimum difference between
that node and all other nodes. Return the minimum of those differences.
cursor - Current node in the traversal.
node - The node being compared. */
func recursiveCompareToTree(cursor *utils.TreeNode, node *utils.TreeNode) int {
	diff := MAX_VAL
	if cursor == nil {
		return diff
	}
	if cursor != node {
		diff = absDiff( cursor.Val, node.Val )
	}
	left_diff := recursiveCompareToTree( cursor.Left, node )
	if left_diff < diff {
		diff = left_diff
	}
	right_diff := recursiveCompareToTree( cursor.Right, node )
	if right_diff < diff {
		diff = right_diff
	}
	return diff
}

func absDiff( x, y int ) int {
	if x > y {
		return x - y
	}
	return y - x
}
