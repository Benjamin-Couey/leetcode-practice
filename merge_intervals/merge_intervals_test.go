package merge_intervals

import (
	"leetcode/utils"
	"testing"
)

func TestMerge(t *testing.T) {
	testcases := []struct {
		intervals [][]int
		solution  [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{[][]int{{2, 6}, {1, 3}, {8, 10}, {15, 18}},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{[][]int{{1, 4}, {4, 4}}, [][]int{{1, 4}}},
		{[][]int{{1, 3}}, [][]int{{1, 3}}},
	}

	for _, testcase := range testcases {
		result := Merge(testcase.intervals)

		if !utils.MatrixEqual(result, testcase.solution) {
			t.Errorf(
				"Merge: %v returned %v, want %v",
				testcase.intervals,
				result,
				testcase.solution,
			)
		}
	}
}
