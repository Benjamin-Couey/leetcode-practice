package path_sum

import (
	"leetcode/utils"
	"testing"
)

func TestHasPathSum(t *testing.T) {
	testcases := []struct {
		tree_list  []int
		target_sum int
		solution   bool
	}{
		{[]int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 0, 1}, 22, true},
		{[]int{1, 2, 3}, 5, false},
		{[]int{}, 0, false},
	}

	for _, testcase := range testcases {
		root := utils.LevelOrderToTree(testcase.tree_list)
		has_sum := HasPathSum(root, testcase.target_sum)

		if has_sum != testcase.solution {
			t.Errorf(
				"HasPathSum: for tree %v and target sum %v, returned %v, want %v",
				testcase.tree_list,
				testcase.target_sum,
				has_sum,
				testcase.solution,
			)
		}
	}
}
