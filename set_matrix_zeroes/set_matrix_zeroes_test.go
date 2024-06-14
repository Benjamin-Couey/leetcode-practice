package set_matrix_zeroes

import (
	"testing"
	"leetcode/utils"
)

func TestSetZeroes(t *testing.T) {
	testcases := []struct {
		matrix [][]int
		solution [][]int
	}{
		{ [][]int{
				{ 1, 1, 1, },
				{ 1, 0, 1, },
				{ 1, 1, 1, },
			}, [][]int{
				{ 1, 0, 1, },
				{ 0, 0, 0, },
				{ 1, 0, 1, },
			},
		},
		{ [][]int{
				{ 0, 1, 2, 0, },
				{ 3, 4, 5, 2, },
				{ 1, 3, 1, 5, },
			}, [][]int{
				{ 0, 0, 0, 0, },
				{ 0, 4, 5, 0, },
				{ 0, 3, 1, 0, },
			},
		},
		{ [][]int{
				{ 4, 1, 2, 1, },
				{ 3, 0, 5, 2, },
				{ 1, 3, 0, 5, },
			}, [][]int{
				{ 4, 0, 0, 1, },
				{ 0, 0, 0, 0, },
				{ 0, 0, 0, 0, },
			},
		},
		{ [][]int{
				{ 1, },
			}, [][]int{
				{ 1, },
			},
		},
	}

	for _, testcase := range testcases {
		SetZeroes( testcase.matrix )

		if !utils.MatrixEqual( testcase.matrix, testcase.solution ) {
			t.Errorf(
				"SetZeroes: returned %v, want %v",
				testcase.matrix,
				testcase.solution,
			)
		}
	}
}
