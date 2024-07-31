package verify_binary_search_tree

import (
	"leetcode/utils"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	testcases := []struct {
		tree_list []int
		solution  bool
	}{
		{[]int{2, 1, 3}, true},
		{[]int{3, 2, 4, 1, 0, 0, 5}, true},
		{[]int{5, 1, 4, 0, 0, 3, 6}, false},
		{[]int{3, 1, 4, 0, 0, 2, 5}, false},
	}

	for _, testcase := range testcases {
		root := utils.LevelOrderToTree(testcase.tree_list)
		result := IsValidBST(root)

		if result != testcase.solution {
			t.Errorf(
				"IsValidBST: for tree %v, returned %v, want %v",
				testcase.tree_list,
				result,
				testcase.solution,
			)
		}
	}
}

func TestAltIsValidBST(t *testing.T) {
	testcases := []struct {
		tree_list []int
		solution  bool
	}{
		{[]int{2, 1, 3}, true},
		{[]int{3, 2, 4, 1, 0, 0, 5}, true},
		{[]int{5, 1, 4, 0, 0, 3, 6}, false},
		{[]int{3, 1, 4, 0, 0, 2, 5}, false},
		{[]int{}, false},
	}

	for _, testcase := range testcases {
		root := utils.LevelOrderToTree(testcase.tree_list)
		result := AltIsValidBST(root)

		if result != testcase.solution {
			t.Errorf(
				"AltIsValidBST: for tree %v, returned %v, want %v",
				testcase.tree_list,
				result,
				testcase.solution,
			)
		}
	}
}
