package convert_sorted_array_to_BST

import (
	"leetcode/is_height_balanced"
	"leetcode/utils"
	"leetcode/verify_binary_search_tree"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	testcases := []struct {
		nums []int
	}{
		{[]int{-10, -3, 0, 5, 9}},
		{[]int{3, 7, 9, 15, 20}},
		{[]int{1, 3}},
		{[]int{1}},
	}

	for _, testcase := range testcases {
		result := SortedArrayToBST(testcase.nums)
		if !utils.SliceInTree(testcase.nums, result) {
			t.Errorf("SortedArrayToBST: %v returned tree with mismatched values", testcase.nums)
		}
		if !is_height_balanced.IsHeightBalanced(result) {
			t.Errorf("SortedArrayToBST: %v did not return height balanced tree", testcase.nums)
		}
		if !verify_binary_search_tree.AltIsValidBST(result) {
			t.Errorf("SortedArrayToBST: %v did not return binary search tree", testcase.nums)
		}
	}
}
