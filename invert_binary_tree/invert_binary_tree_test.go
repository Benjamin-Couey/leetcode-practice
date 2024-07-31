package invert_binary_tree

import (
	"leetcode/utils"
	"testing"
)

func TestInvertTree(t *testing.T) {
	testcases := []struct {
		base_list, inverted_list []int
	}{
		{[]int{4, 2, 7, 1, 3, 6, 9}, []int{4, 7, 2, 9, 6, 3, 1}},
		{[]int{2, 1, 3}, []int{2, 3, 1}},
		{[]int{}, []int{}},
	}

	for _, testcase := range testcases {
		base_root := utils.LevelOrderToTree(testcase.base_list)
		utils.TraverseAndPrintTree(base_root)
		inverted_root := InvertTree(base_root)

		utils.TraverseAndPrintTree(inverted_root)

		same := utils.IsSameTree(
			inverted_root,
			utils.LevelOrderToTree(testcase.inverted_list),
		)

		if !same {
			t.Errorf(
				"InvertTree: did not produce inverted %v, want %v",
				testcase.base_list,
				testcase.inverted_list,
			)
		}
	}
}
