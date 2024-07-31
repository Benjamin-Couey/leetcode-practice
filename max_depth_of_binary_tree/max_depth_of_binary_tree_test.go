package max_depth_of_binary_tree

import (
	"leetcode/utils"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	testcases := []struct {
		node_list []int
		solution  int
	}{
		{[]int{3, 9, 20, 0, 0, 15, 7}, 3},
		{[]int{1, 0, 2}, 2},
		{[]int{1}, 1},
		{[]int{}, 0},
	}

	for _, testcase := range testcases {
		depth := MaxDepth(utils.LevelOrderToTree(testcase.node_list))

		if depth != testcase.solution {
			t.Errorf("MaxDepth: %v, want %v", depth, testcase.solution)
		}
	}
}
