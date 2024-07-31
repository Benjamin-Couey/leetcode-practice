package spiral_matrix

import (
	"leetcode/utils"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	testcases := []struct {
		matrix   [][]int
		solution []int
	}{
		{[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{[][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	}

	for _, testcase := range testcases {
		result := SpiralOrder(testcase.matrix)

		if !utils.SliceEqual(result, testcase.solution) {
			t.Errorf(
				"SpiralOrder: %v returned %v, want %v",
				testcase.matrix,
				result,
				testcase.solution,
			)
		}
	}
}
