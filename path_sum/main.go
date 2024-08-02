/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/path-sum/description/
*/
package path_sum

import (
	"leetcode/utils"
)

/*
HasPathSum reports whether the tree starting at root has a root to leaf sum
equally targetSum.
HasPathSum Assumes that:
The number of nodes in the tree is in the range [0, 5000],
-1000 <= Node.Val <= 1000 and Node.Val != 0,
and -1000 <= targetSum <= 1000.
*/
func HasPathSum(root *utils.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return recursiveHasPathSum(root, 0, targetSum)
}

func recursiveHasPathSum(node *utils.TreeNode, currentSum int, targetSum int) bool {
	if node == nil {
		if currentSum == targetSum {
			return true
		}
		return false
	}

	currentSum += node.Val

	return recursiveHasPathSum(node.Left, currentSum, targetSum) || recursiveHasPathSum(node.Right, currentSum, targetSum)
}
