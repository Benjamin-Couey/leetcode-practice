package symmetric_tree

import (
	"leetcode/utils"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	testcases := []struct {
		tree_list []int
		solution  bool
	}{
		{[]int{1, 2, 2, 3, 4, 4, 3}, true},
		{[]int{1, 2, 2, 3, 4, 3, 4}, false},
		{[]int{1, 2, 2, 0, 3, 0, 3}, false},
	}

	for _, testcase := range testcases {
		root := utils.LevelOrderToTree(testcase.tree_list)
		symmetric := IsSymmetric(root)

		if symmetric != testcase.solution {
			t.Errorf(
				"IsSymmetric: for tree %v returned %v, want %v",
				testcase.tree_list,
				symmetric,
				testcase.solution,
			)
		}
	}
}
