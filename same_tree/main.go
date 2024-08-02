/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/same-tree/description/
*/
package same_tree

import (
	"leetcode/utils"
)

/*
IsSameTree reports whether p and q are the same tree. That is, they are structurally
identical and all nodes have the same value.
IsSameTree assumes that:
the number of nodes in both trees is in the range [0, 100],
and -10^4 <= Node.val <= 10^4 and Node.val != 0.
*/
func IsSameTree(p *utils.TreeNode, q *utils.TreeNode) bool {
	return utils.IsSameTree(p, q)
}
