/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/count-complete-tree-nodes/description/
*/
package count_complete_tree_nodes

import (
	"leetcode/utils"
)

/*
CountNodes returns number of nodes in tree starting at root.
CountNodes assumes that:
the number of nodes in the tree is in the range [0, 5 * 10^4],
0 <= Node.val <= 5 * 10^4,
and the tree is guaranteed to be complete.
*/
func CountNodes(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	count := 1
	count += CountNodes(root.Left)
	count += CountNodes(root.Right)
	return count
}
