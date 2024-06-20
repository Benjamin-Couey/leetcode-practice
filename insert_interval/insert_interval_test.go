package insert_interval

import (
	"testing"
	"leetcode/utils"
)

func TestInsert(t *testing.T) {
	testcases := []struct {
		intervals [][]int
		newInterval []int
		solution [][]int
	}{
		{ [][]int{ { 1, 3 }, { 6, 9 }, }, []int{ 2, 5 }, [][]int{ { 1, 5 }, { 6, 9, }, }, },
		{ [][]int{ { 6, 9 }, }, []int{ 1, 3 }, [][]int{ { 1, 3 }, { 6, 9, }, }, },
		{ [][]int{ { 1, 2 }, { 3, 5 }, { 6, 7 }, { 8, 10 }, { 12, 16 } },
			[]int{ 4, 8 },
			[][]int{ { 1, 2 }, { 3, 10, }, { 12, 16 } },
		},
		{ [][]int{ { 1, 3 }, }, []int{ 1, 3 }, [][]int{ { 1, 3 }, }, },
		{ [][]int{}, []int{ 1, 3 }, [][]int{ { 1, 3 }, }, },
	}

	for _, testcase := range testcases {
		result := Insert( testcase.intervals, testcase.newInterval )

		if !utils.MatrixEqual( result, testcase.solution ) {
			t.Errorf(
				"Insert: %v and %v returned %v, want %v",
				testcase.intervals,
				testcase.newInterval,
				result,
				testcase.solution,
			)
		}
	}
}
