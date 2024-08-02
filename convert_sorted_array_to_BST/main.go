/*
Package includes implementation and tests for problem described at
https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/
*/
package convert_sorted_array_to_BST

import (
	"leetcode/utils"
)

/*
SortedArrayToBST returns root of a new height balanced binary search tree
containing all values in nums.
SortedArrayToBST assumes that:
1 <= nums.length <= 10^4,
-10^4 <= nums[i] <= 10^4,
and nums is sorted in a strictly increasing order.
*/
func SortedArrayToBST(nums []int) *utils.TreeNode {
	if len(nums) < 1 {
		return nil
	}

	index := len(nums) / 2

	root := utils.TreeNode{nums[index], nil, nil}
	root.Left = SortedArrayToBST(nums[:index])
	root.Right = SortedArrayToBST(nums[index+1:])

	return &root
}
