package count_complete_tree_nodes

import (
	"leetcode/utils"
	"testing"
)

func TestCountNodes(t *testing.T) {
	testcases := []struct {
		tree_list []int
		solution  int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 6},
		{[]int{}, 0},
		{[]int{1}, 1},
	}

	for _, testcase := range testcases {
		result := CountNodes(utils.LevelOrderToTree(testcase.tree_list))

		if result != testcase.solution {
			t.Errorf(
				"CountNodes: %v returned %v, want %v",
				testcase.tree_list,
				result,
				testcase.solution,
			)
		}
	}
}
