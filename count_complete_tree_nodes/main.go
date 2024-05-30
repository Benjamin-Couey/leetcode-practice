package count_complete_tree_nodes

import (
	"leetcode/utils"
)

func CountNodes(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	count := 1
	count += CountNodes( root.Left )
	count += CountNodes( root.Right )
	return count
}
