/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/validate-binary-search-tree/description/
*/
package verify_binary_search_tree

import (
	"leetcode/utils"
)

/*
IsValidBST reports whether the tree starting at root is a valid binary search
tree. A valid BST is defined as follows:
the left subtree of a node contains only nodes with values less than the node's value,
the right subtree of a node contains only nodes with values greater than the node's value,
and both the left and right subtrees must also be binary search trees.
IsValidBST assumes that:
the number of nodes in the tree is in the range [1, 10^4],
and -2^31 <= Node.val <= 2^31 - 1.
*/
func IsValidBST(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}

	if !subtreeLessThan(root.Left, root.Val) || !subtreeGreaterThan(root.Right, root.Val) {
		return false
	} else {
		return IsValidBST(root.Left) && IsValidBST(root.Right)
	}
}

func subtreeLessThan(root *utils.TreeNode, val int) bool {
	if root == nil {
		return true
	}

	if root.Val >= val {
		return false
	} else {
		return subtreeLessThan(root.Left, val) && subtreeLessThan(root.Right, val)
	}
}

func subtreeGreaterThan(root *utils.TreeNode, val int) bool {
	if root == nil {
		return true
	}

	if root.Val <= val {
		return false
	} else {
		return subtreeGreaterThan(root.Left, val) && subtreeGreaterThan(root.Right, val)
	}
}

/*
AltIsValidBST is an alternative implementation of IsValidBST which makes
the additional assumption that Node.val != 0.
*/
func AltIsValidBST(root *utils.TreeNode) bool {
	return recAltIsValidBST(root, 0, 0)
}

func recAltIsValidBST(root *utils.TreeNode, lower_bound int, upper_bound int) bool {
	if root == nil {
		return false
	}

	if lower_bound != 0 && root.Val < lower_bound {
		return false
	}

	if upper_bound != 0 && root.Val > upper_bound {
		return false
	}

	left_good := true
	if root.Left != nil {
		left_good = recAltIsValidBST(root.Left, lower_bound, root.Val)
	}

	right_good := true
	if root.Right != nil {
		right_good = recAltIsValidBST(root.Right, root.Val, upper_bound)
	}

	return left_good && right_good
}
