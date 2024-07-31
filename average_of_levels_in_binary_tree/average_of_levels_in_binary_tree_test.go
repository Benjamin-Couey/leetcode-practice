package average_of_levels_in_binary_tree

import (
	"leetcode/utils"
	"testing"
)

func TestAverageOfLevels(t *testing.T) {
	testcases := []struct {
		tree_list []int
		solution  []float64
	}{
		{[]int{3, 9, 20, 0, 0, 15, 7}, []float64{3, 14.5, 11}},
		{[]int{3, 9, 20, 15, 7}, []float64{3, 14.5, 11}},
		{[]int{1}, []float64{1}},
	}

	for _, testcase := range testcases {
		result := AverageOfLevels(utils.LevelOrderToTree(testcase.tree_list))

		if !utils.SliceEqual(result, testcase.solution) {
			t.Errorf(
				"AverageOfLevels: %v returned %v, want %v",
				testcase.tree_list,
				result,
				testcase.solution,
			)
		}
	}
}
