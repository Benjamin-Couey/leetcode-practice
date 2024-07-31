package symmetric_tree

import (
	"leetcode/utils"
)

/* Assumes that:
The number of nodes in the tree is in the range [1, 1000].
-100 <= Node.Val <= 100 and Node.Val != 0 */

func IsSymmetric(root *utils.TreeNode) bool {
	return recursiveIsSymmetric(root.Left, root.Right)
}

func recursiveIsSymmetric(leftCursor, rightCursor *utils.TreeNode) bool {
	if leftCursor == nil && rightCursor == nil {
		return true
	} else if leftCursor == nil || rightCursor == nil && leftCursor != rightCursor {
		return false
	} else if leftCursor.Val != rightCursor.Val {
		return false
	}

	return recursiveIsSymmetric(leftCursor.Left, rightCursor.Right) && recursiveIsSymmetric(leftCursor.Right, rightCursor.Left)
}
